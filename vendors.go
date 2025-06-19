package hwarchiver

import (
	"fmt"
	"strings"
)

type Vendor struct {
	Identity
	Devices map[int64]*Device
}

func NewVendor() *Vendor {
	return &Vendor{
		Identity: Identity{
			ID:   -1,
			Name: "",
		},
		Devices: make(map[int64]*Device),
	}
}

func (c *Vendor) addDevice(device *Device) error {
	return AddToMap(c.Devices, device.ID, device, "HWArchive.Vendors.Devices")
}

func parseVendorLine(s string) (*Vendor, error) {
	newClass := NewVendor()

	var err error
	newClass.Name, err = parseHexFieldsLine(s, &newClass.ID)
	if err != nil {
		return nil, err
	}

	return newClass, nil
}

func readVendorSection(exp *HWExplorer, hwa *HWArchive) error {
	return readTopLevelSection[*Vendor, *Device](
		exp,
		hwa,
		parseVendorLine,
		func(hwa *HWArchive, v *Vendor) {
			hwa.addVendor(v)
		},
		func(v *Vendor) error {
			return readSection(
				exp,
				v,
				parseDeviceLine,
				func(parent *Vendor, child *Device) error {
					parent.addDevice(child)
					return nil
				},
				func(line string) bool {
					return strings.HasPrefix(line, "C ") || lineStartsWithHex(line) ||
						(strings.HasPrefix(line, "\t") && !strings.HasPrefix(line, "\t\t"))
				},
				func(line string) bool {
					return strings.HasPrefix(line, "\t\t")
				},
				func(dev *Device, line string) error {
					subDev, err := parseSubdeviceLine(line)
					if err != nil {
						return err
					}
					fmt.Printf("↓>>>>>>>> 0x%04x 0x%04x %s\n", subDev.ID, subDev.Subdevice, subDev.Name)
					dev.addSubdevice(subDev)
					return nil
				},
				func(dev *Device) {
					fmt.Printf("↓>>>> 0x%04x %s\n", dev.ID, dev.Name)
				},
				func(_ *Device) {},
			)
		},
		func(v *Vendor) {
			fmt.Printf("↓ 0x%04x %s\n", v.ID, v.Name)
		},
	)
}
