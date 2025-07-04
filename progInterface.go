package hwarchiver

import (
	"bufio"
	"log"
	"strings"
)

type ProgInterface struct {
	Identity `yaml:",inline"`
}

func newProgInterface() *ProgInterface {
	return &ProgInterface{
		Identity: Identity{
			ID:   -1,
			Name: "",
		},
	}
}

func parseProgInterfaceLine(s string) (*ProgInterface, error) {
	newSubclass := newProgInterface()

	var err error
	newSubclass.Name, err = parseHexFieldsLine(s, &newSubclass.ID)
	if err != nil {
		return nil, err
	}

	return newSubclass, nil
}

func parseProginterface(block string, sc *Subclass) {
	scanner := bufio.NewScanner(strings.NewReader(block))

	for scanner.Scan() {
		progIf, err := parseProgInterfaceLine(scanner.Text())
		if err != nil {
			log.Fatalf("failed to parse line: `%s`, %v", scanner.Text(), err)
		}

		err = sc.addProgIF(progIf)
		if err != nil {
			log.Fatalf("failed to add stuff to map: %v", err)
		}
	}
}
