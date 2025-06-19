package hwarchiver

type Subclass struct {
	Identity `yaml:",inline"`
	ProgInterfaces map[int64]*ProgInterface `yaml:"prog_interfaces"`
}

func NewSubclass() *Subclass {
	return &Subclass{
		Identity: Identity{
			ID:   -1,
			Name: "",
		},
		ProgInterfaces: make(map[int64]*ProgInterface),
	}
}

func (s *Subclass) addProgIF(progIf *ProgInterface) error {
	return AddToMap(s.ProgInterfaces, progIf.ID, progIf,
		"HWArchive.Classes.Subclasses.ProgInterfaces")
}

func parseSubclassLine(s string) (*Subclass, error) {
	newSubclass := NewSubclass()

	var err error
	newSubclass.Name, err = parseHexFieldsLine(s, &newSubclass.ID)
	if err != nil {
		return nil, err
	}

	return newSubclass, nil
}
