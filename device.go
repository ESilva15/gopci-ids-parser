package hwarchiver

type Device struct {
	Identity   `yaml:",inline"`
	Subdevices map[SubdeviceKey]*Subdevice `yaml:"subdevices"`
}

func newDevice() *Device {
	return &Device{
		Identity: Identity{
			ID:   -1,
			Name: "",
		},
		Subdevices: make(map[SubdeviceKey]*Subdevice),
	}
}

func (d *Device) addSubdevice(device *Subdevice) error {
	key := SubdeviceKey{device.ID, device.Subdevice}
	return addToMap(d.Subdevices, key, device, "HWArchive.Vendors.Device.Subdevices")
}

func parseDeviceLine(s string) (*Device, error) {
	newClass := newDevice()

	var err error
	newClass.Name, err = parseHexFieldsLine(s, &newClass.ID)
	if err != nil {
		return nil, err
	}

	return newClass, nil
}
