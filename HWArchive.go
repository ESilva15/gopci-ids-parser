package hwarchiver

import (
	"bufio"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Named interface {
	SetName(string)
}

type Identity struct {
	ID   int64 	`yaml:"id"`
	Name string `yaml:"name"`
}

func (i *Identity) SetName(name string) {
	i.Name = name
}

type HWArchive struct {
	Vendors map[int64]*Vendor `yaml:"vendors"`
	Classes map[int64]*Class `yaml:"classes"`
}

func CreateHWArchive() *HWArchive {
	return &HWArchive{
		Vendors: make(map[int64]*Vendor),
		Classes: make(map[int64]*Class),
	}
}

func (hwa *HWArchive) addVendor(cls *Vendor) error {
	return AddToMap(hwa.Vendors, cls.ID, cls, "HWArchive.Vendors")
}

func (hwa *HWArchive) addClass(cls *Class) error {
	return AddToMap(hwa.Classes, cls.ID, cls, "HWArchive.Classes")
}

func (hwa *HWArchive) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	explorer := NewHWExplcorer(scanner)

	for explorer.Scan() {
		line := explorer.Peek()

		// Ignore the comments
		if strings.HasPrefix(line, "#") {
			explorer.Consume()
			continue
		}

		// Handle the Classes
		if strings.HasPrefix(line, "C ") {
			err := readClassSection(explorer, hwa)
			if err != nil {
				return err
			}

			continue
		}

		// Handle the vendors
		if lineStartsWithHex(line) {
			// Its a vendor
			readVendorSection(explorer, hwa)
			// hwa.Explorer.Consume()
			continue
		}

		explorer.Consume()
	}

	return nil
}

func (hwa *HWArchive) ToYAML() (string, error) {
	yamlData, err := yaml.Marshal(hwa)
	if err != nil {
		return "", err
	}

	return string(yamlData), nil
}
