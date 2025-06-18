package hwarchiver

type ProgInterface struct {
	ID   int64
	Name string
}

func parseProgInterfaceLine(s string) (*ProgInterface, error) {
	newSubclass := ProgInterface{-1, ""}

	err := parseHexStringLine(&newSubclass.ID, &newSubclass.Name, s)
	if err != nil {
		return nil, err
	}

	return &newSubclass, nil
}
