package hwarchiver

type Subdevice struct {
	Identity
	Subdevice int64
}

type SubdeviceKey struct {
	SubvendorID int64
	SubdeviceID int64
}

func NewSubdevice() *Subdevice {
	return &Subdevice{
		Identity: Identity{
			ID:   -1,
			Name: "",
		},
		Subdevice: -1,
	}
}

func parseSubdeviceLine(s string) (*Subdevice, error) {
	newSubdev := NewSubdevice()

	var err error
	newSubdev.Name, err = parseHexFieldsLine(s, &newSubdev.ID, &newSubdev.Subdevice)
	if err != nil {
		return nil, err
	}

	return newSubdev, nil
}
