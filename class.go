package hwarchiver

import (
	"fmt"
	"strings"
)

type Class struct {
	Identity `yaml:",inline"`
	Subclasses map[int64]*Subclass `yaml:"subclasses"`
}

func NewClass() *Class {
	return &Class{
		Identity: Identity{
			ID:   -1,
			Name: "",
		},
		Subclasses: make(map[int64]*Subclass),
	}
}

func (c *Class) addSubclass(cls *Subclass) error {
	return AddToMap(c.Subclasses, cls.ID, cls, "HWArchive.Classes.Subclasses")
}

func parseClassLine(s string) (*Class, error) {
	input := s[2:]
	newClass := NewClass()

	var err error
	newClass.Name, err = parseHexFieldsLine(input, &newClass.ID)
	if err != nil {
		return nil, err
	}

	return newClass, nil
}

func readClassSection(exp *HWExplorer, hwa *HWArchive) error {
	return readTopLevelSection[*Class, *Subclass](
		exp,
		hwa,
		parseClassLine,
		func(hwa *HWArchive, c *Class) {
			hwa.addClass(c)
		},
		func(c *Class) error {
			return readSection(
				exp,
				c,
				parseSubclassLine,
				func(parent *Class, child *Subclass) error {
					parent.addSubclass(child)
					return nil
				},
				func(line string) bool {
					return strings.HasPrefix(line, "C ") || lineStartsWithHex(line) ||
						(strings.HasPrefix(line, "\t") && !strings.HasPrefix(line, "\t\t"))
				},
				func(line string) bool {
					return strings.HasPrefix(line, "\t\t")
				},
				func(sc *Subclass, line string) error {
					progIF, err := parseProgInterfaceLine(line)
					if err != nil {
						return err
					}
					fmt.Printf("|-------- 0x%04x %s\n", progIF.ID, progIF.Name)
					sc.addProgIF(progIF)
					return nil
				},
				func(sc *Subclass) {
					fmt.Printf("|---- 0x%04x %s\n", sc.ID, sc.Name)
				},
				func(_ *Subclass) {},
			)
		},
		func(c *Class) {
			fmt.Printf("| 0x%04x %s\n", c.ID, c.Name)
		},
	)
}
