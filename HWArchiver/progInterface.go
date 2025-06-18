package hwarchiver

type ProgInterface struct {
	ID   int64
	Name string
}

func parseProgInterfaceLine(s string) (*ProgInterface, error) {
	newSubclass := ProgInterface{-1, ""}

	var err error
	newSubclass.Name, err = parseHexFieldsLine(s, &newSubclass.ID)
	if err != nil {
		return nil, err
	}

	return &newSubclass, nil
}
