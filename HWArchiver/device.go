package hwarchiver

import (
	"fmt"
	"strings"
)

type Device struct {
	ID         int64
	Name       string
	Subdevices map[SubdeviceKey]*Subdevice
}

func (c *Device) addSubdevice(device *Subdevice) error {
	if c.Subdevices == nil {
		return fmt.Errorf("HWArchive.Classes.Subclasses.ProgInterfaces shouldn't be nill")
	}

	key := SubdeviceKey{device.ID, device.Subdevice}

	_, ok := c.Subdevices[key]
	if ok {
		return fmt.Errorf("Key `%d` is already present", device.ID)
	}

	c.Subdevices[key] = device

	return nil
}

func parseDeviceLine(s string) (*Device, error) {
	newClass := Device{-1, "", make(map[SubdeviceKey]*Subdevice)}

	var err error
	newClass.Name, err = parseHexFieldsLine(s, &newClass.ID)
	if err != nil {
		return nil, err
	}

	return &newClass, nil
}

func readDeviceSection(exp *HWExplorer, vendor *Vendor) error {
	// This consumes the class id and name
	newDevice, err := parseDeviceLine(strings.TrimSpace(exp.Peek()))
	if err != nil {
		return err
	}

	fmt.Printf("↓>>>> 0x%04x %s\n", newDevice.ID, newDevice.Name)
	vendor.addDevice(newDevice)

	exp.Consume()

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
			subdevice, err := parseSubdeviceLine(strings.TrimSpace(exp.Peek()))
			if err != nil {
				return err
			}

			newDevice.addSubdevice(subdevice)

			fmt.Printf("↓>>>>>>>> 0x%04x 0x%04x %s\n", subdevice.ID, subdevice.Subdevice,
				subdevice.Name)
			exp.Consume()
			continue
		}

		exp.Consume()
	}

	return nil
}
