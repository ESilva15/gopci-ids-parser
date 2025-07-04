package hwarchiver

import (
	"bufio"
	"log"
	"strings"
)

type Subdevice struct {
	Identity  `yaml:",inline"`
	Subdevice int64 `yaml:"subdevice_id"`
}

type SubdeviceKey struct {
	SubvendorID int64
	SubdeviceID int64
}

func newSubdevice() *Subdevice {
	return &Subdevice{
		Identity: Identity{
			ID:   -1,
			Name: "",
		},
		Subdevice: -1,
	}
}

func parseSubdeviceLine(s string) (*Subdevice, error) {
	newSubdev := newSubdevice()

	var err error
	newSubdev.Name, err = parseHexFieldsLine(s, &newSubdev.ID, &newSubdev.Subdevice)
	if err != nil {
		return nil, err
	}

	return newSubdev, nil
}

func parseSubdevice(block string, d *Device) {
	scanner := bufio.NewScanner(strings.NewReader(block))

	for scanner.Scan() {
		subdevice, err := parseSubdeviceLine(scanner.Text())
		if err != nil {
			log.Fatalf("failed to parse line: `%s`, %v", scanner.Text(), err)
		}

		d.addSubdevice(subdevice)
	}
}
