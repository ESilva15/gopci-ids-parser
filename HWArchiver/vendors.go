package hwarchiver

import (
	"fmt"
	"strings"
)

type Vendor struct {
	ID      int64
	Name    string
	Devices map[int64]*Device
}

func (c *Vendor) addDevice(device *Device) error {
	if c.Devices == nil {
		return fmt.Errorf("HWArchive.Classes.Subclasses.ProgInterfaces shouldn't be nill")
	}

	_, ok := c.Devices[device.ID]
	if ok {
		return fmt.Errorf("Key `%d` is already present", device.ID)
	}

	c.Devices[device.ID] = device

	return nil
}

func parseVendorLine(s string) (*Vendor, error) {
	newClass := Vendor{-1, "", make(map[int64]*Device)}

	err := parseHexStringLine(&newClass.ID, &newClass.Name, s)
	if err != nil {
		return nil, err
	}

	return &newClass, nil
}

func readVendorSection(exp *HWExplorer, hwa *HWArchive) error {
	// This consumes the class id and name
	newVendor, err := parseVendorLine(exp.Peek())
	if err != nil {
		return err
	}

	fmt.Printf("↓ 0x%04x %s\n", newVendor.ID, newVendor.Name)
	hwa.addVendor(newVendor)

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
			_ = readDeviceSection(exp, newVendor)
			// exp.Consume()
			// if err != nil {
			// 	return err
			// }
			continue
		}

		exp.Consume()
	}

	return nil
}
