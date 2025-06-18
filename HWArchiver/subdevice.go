package hwarchiver

type Subdevice struct {
	ID        int64
	Subdevice int64
	Name      string
}

type SubdeviceKey struct {
	SubvendorID int64
	SubdeviceID int64
}

func parseSubdeviceLine(s string) (*Subdevice, error) {
	newSubdev := &Subdevice{}
	err := parseHexHexStringLine(&newSubdev.ID, &newSubdev.Subdevice,
		&newSubdev.Name, s)
	if err != nil {
		return nil, err
	}

	return newSubdev, nil
}
