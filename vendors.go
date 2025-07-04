package hwarchiver

import (
	"bufio"
	"log"
	"strings"
)

type Vendor struct {
	Identity `yaml:",inline"`
	Devices  map[int64]*Device `yaml:"devices"`
}

func newVendor() *Vendor {
	return &Vendor{
		Identity: Identity{
			ID:   -1,
			Name: "",
		},
		Devices: make(map[int64]*Device),
	}
}

func (c *Vendor) addDevice(device *Device) error {
	return addToMap(c.Devices, device.ID, device, "HWArchive.Vendors.Devices")
}

func parseVendorLine(s string) (*Vendor, error) {
	newClass := newVendor()

	var err error
	newClass.Name, err = parseHexFieldsLine(s, &newClass.ID)
	if err != nil {
		return nil, err
	}

	return newClass, nil
}

func parseVendorBlock(vendorStr string, block string, hwa *HWArchive) {
	vendor, err := parseVendorLine(vendorStr)
	if err != nil {
		log.Fatalf("Failed to parse class `%s`: %v", vendorStr, err)
	}
	err = hwa.addVendor(vendor)
	if err != nil {
		log.Fatalf("failed to add class to classes: %v", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(block))
	explorer := newHWExplorer(scanner)

	for explorer.scan() {
		line := explorer.peek()

		if strings.HasPrefix(line, "\t") {
			device, err := parseDeviceLine(line)
			if err != nil {
				log.Fatalf("Failed to parse subclass `%s`: %v", line, err)
			}
			err = vendor.addDevice(device)
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

			parseSubdevice(block, device)

			continue
		}
	}
}
