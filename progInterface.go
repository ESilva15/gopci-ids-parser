package hwarchiver

type ProgInterface struct {
	Identity `yaml:",inline"`
}

func NewProgInterface() *ProgInterface {
	return &ProgInterface{
		Identity: Identity{
			ID:   -1,
			Name: "",
		},
	}
}

func parseProgInterfaceLine(s string) (*ProgInterface, error) {
	newSubclass := NewProgInterface()

	var err error
	newSubclass.Name, err = parseHexFieldsLine(s, &newSubclass.ID)
	if err != nil {
		return nil, err
	}

	return newSubclass, nil
}
