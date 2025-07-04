package hwarchiver

import (
	"bufio"
	"log"
	"strings"
)

type Class struct {
	Identity   `yaml:",inline"`
	Subclasses map[int64]*Subclass `yaml:"subclasses"`
}

func newClass() *Class {
	return &Class{
		Identity: Identity{
			ID:   -1,
			Name: "",
		},
		Subclasses: make(map[int64]*Subclass),
	}
}

func (c *Class) addSubclass(cls *Subclass) error {
	return addToMap(c.Subclasses, cls.ID, cls, "HWArchive.Classes.Subclasses")
}

func parseClassLine(s string) (*Class, error) {
	input := s[2:]
	newClass := newClass()

	var err error
	newClass.Name, err = parseHexFieldsLine(input, &newClass.ID)
	if err != nil {
		return nil, err
	}

	return newClass, nil
}

func parseClassBlock(classStr string, block string, hwa *HWArchive) {
	class, err := parseClassLine(classStr)
	if err != nil {
		log.Fatalf("Failed to parse class `%s`: %v", classStr, err)
	}
	err = hwa.addClass(class)
	if err != nil {
		log.Fatalf("failed to add class to classes: %v", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(block))
	explorer := newHWExplorer(scanner)

	for explorer.scan() {
		line := explorer.peek()

		if strings.HasPrefix(line, "\t") {
			subclass, err := parseSubclassLine(line)
			if err != nil {
				log.Fatalf("Failed to parse subclass `%s`: %v", line, err)
			}

			err = class.addSubclass(subclass)
			if err != nil {
				log.Fatalf("failed to add class to classes: %v", err)
			}
			explorer.consume()

			prefixChecker := func(line string) bool {
				return strings.HasPrefix(line, "\t\t")
			}
			block, err := readBlock(explorer, prefixChecker)
			if err != nil {
				log.Fatalf("It broke here: %v", err)
			}

			parseProginterface(block, subclass)

			continue
		}
	}
}
