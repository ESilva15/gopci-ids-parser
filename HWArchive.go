package hwarchiver

import (
	"bufio"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Named interface {
	setName(string)
}

type Identity struct {
	ID   int64  `yaml:"id"`
	Name string `yaml:"name"`
}

func (i *Identity) setName(name string) {
	i.Name = name
}

type HWArchive struct {
	Vendors map[int64]*Vendor `yaml:"vendors"`
	Classes map[int64]*Class  `yaml:"classes"`
}

func CreateHWArchive() *HWArchive {
	return &HWArchive{
		Vendors: make(map[int64]*Vendor),
		Classes: make(map[int64]*Class),
	}
}

func (hwa *HWArchive) addVendor(cls *Vendor) error {
	return addToMap(hwa.Vendors, cls.ID, cls, "HWArchive.Vendors")
}

func (hwa *HWArchive) addClass(cls *Class) error {
	return addToMap(hwa.Classes, cls.ID, cls, "HWArchive.Classes")
}

func (hwa *HWArchive) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	explorer := newHWExplorer(scanner)

	for explorer.scan() {
		line := explorer.peek()
		if strings.HasPrefix(line, "#") || line == "" {
			explorer.consume()
			continue
		}

		if strings.HasPrefix(line, "C ") {
			prefixChecker := func(line string) bool {
				return strings.HasPrefix(line, "\t")
			}
			explorer.consume()
			block, err := readBlock(explorer, prefixChecker)
			if err != nil {
				log.Fatalf("It broke here: %v", err)
			}

			parseClassBlock(line, block, hwa)
			continue
		}

		if lineStartsWithHex(line) {
			prefixChecker := func(line string) bool {
				return strings.HasPrefix(line, "\t")
			}
			explorer.consume()
			block, err := readBlock(explorer, prefixChecker)
			if err != nil {
				log.Fatalf("It broke here: %v", err)
			}

			parseVendorBlock(line, block, hwa)
			continue
		}
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
