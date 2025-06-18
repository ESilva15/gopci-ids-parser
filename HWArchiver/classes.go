package hwarchiver

import (
	"fmt"
	"strings"
)

type Subclass struct {
	ID             int64
	Name           string
	ProgInterfaces map[int64]*ProgInterface
}

func (c *Subclass) addProgIF(progIf *ProgInterface) error {
	if c.ProgInterfaces == nil {
		return fmt.Errorf("HWArchive.Classes.Subclasses.ProgInterfaces shouldn't be nill")
	}

	_, ok := c.ProgInterfaces[progIf.ID]
	if ok {
		return fmt.Errorf("Key `%d` is already present", progIf.ID)
	}

	c.ProgInterfaces[progIf.ID] = progIf

	return nil
}

type Class struct {
	ID         int64
	Name       string
	Subclasses map[int64]*Subclass
}

func (c *Class) addClass(cls *Subclass) error {
	if c.Subclasses == nil {
		return fmt.Errorf("HWArchive.Classes.Subclasses shouldn't be nill")
	}

	_, ok := c.Subclasses[cls.ID]
	if ok {
		return fmt.Errorf("Key `%d` is already present", cls.ID)
	}

	c.Subclasses[cls.ID] = cls

	return nil
}

func parseClassLine(s string) (*Class, error) {
	input := s[2:]
	newClass := Class{-1, "", make(map[int64]*Subclass)}

	err := parseHexStringLine(&newClass.ID, &newClass.Name, input)
	if err != nil {
		return nil, err
	}

	return &newClass, nil
}

func parseSubclassLine(s string) (*Subclass, error) {
	newSubclass := Subclass{-1, "", make(map[int64]*ProgInterface)}

	err := parseHexStringLine(&newSubclass.ID, &newSubclass.Name, s)
	if err != nil {
		return nil, err
	}

	return &newSubclass, nil
}

func readClassSection(exp *HWExplorer, hwa *HWArchive) error {
	// This consumes the class id and name
	newClass, err := parseClassLine(exp.Peek())
	if err != nil {
		return err
	}

	fmt.Printf("| 0x%04x %s\n", newClass.ID, newClass.Name)
	hwa.addClass(newClass)

	exp.Consume()

	for exp.Scan() {
		line := exp.Peek()

		if strings.HasPrefix(line, "C ") || lineStartsWithHex(line) {
			break
		}

		if strings.HasPrefix(line, "\t\t") {
			return fmt.Errorf("line `%d` went from Class to Prog Interface", exp.line)
		}

		// We found a subclass, need to add it to our current class
		if strings.HasPrefix(line, "\t") {
			err := readSubclassSection(exp, newClass)
			if err != nil {
				return err
			}
			continue
		}

		exp.Consume()
	}

	return nil
}

func readSubclassSection(exp *HWExplorer, cls *Class) error {
	newSubClass, err := parseSubclassLine(strings.TrimSpace(exp.Peek()))
	if err != nil {
		return err
	}

	fmt.Printf("|---- 0x%04x %s\n", newSubClass.ID, newSubClass.Name)
	cls.addClass(newSubClass)

	exp.Consume()

	// This consumes the class id and name
	for exp.Scan() {
		line := exp.Peek()

		if strings.HasPrefix(line, "C ") || lineStartsWithHex(line) {
			break
		}

		// If the next line is another subclass line (\t but not \t\t)
		if strings.HasPrefix(line, "\t") && !strings.HasPrefix(line, "\t\t") {
			break
		}

		if strings.HasPrefix(line, "\t\t") {
			// Got into the prog interfaces
			newProgIf, err := parseProgInterfaceLine(strings.TrimSpace(exp.Peek()))
			if err != nil {
				return err
			}

			newSubClass.addProgIF(newProgIf)

			fmt.Printf("|-------- 0x%04x %s\n", newProgIf.ID, newProgIf.Name)
			exp.Consume()
			continue
		}

		exp.Consume()
	}

	return nil
}

func parseSubClassLine(s string) (*Class, error) {
	// input := s[2:]

	fmt.Println(strings.Fields(s))
	return &Class{0, "", make(map[int64]*Subclass)}, nil
}
