package hwarchiver

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type HWArchive struct {
	Explorer *HWExplorer
	Vendors  map[int64]*Vendor
	Classes  map[int64]*Class
}

func CreateHWArchive() *HWArchive {
	return &HWArchive{
		Explorer: nil,
		Vendors:  make(map[int64]*Vendor),
		Classes:  make(map[int64]*Class),
	}
}

func (hwa *HWArchive) addClass(cls *Class) error {
	if hwa.Classes == nil {
		return fmt.Errorf("HWArchive.Classes shouldn't be nill")
	}

	_, ok := hwa.Classes[cls.ID]
	if ok {
		return fmt.Errorf("Key `%d` is already present", cls.ID)
	}

	hwa.Classes[cls.ID] = cls

	return nil
}

func (hwa *HWArchive) addVendor(cls *Vendor) error {
	if hwa.Classes == nil {
		return fmt.Errorf("HWArchive.Classes shouldn't be nill")
	}

	_, ok := hwa.Classes[cls.ID]
	if ok {
		return fmt.Errorf("Key `%d` is already present", cls.ID)
	}

	hwa.Vendors[cls.ID] = cls

	return nil
}

func (hwa *HWArchive) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hwa.Explorer = NewHWExplcorer(scanner)

	for hwa.Explorer.Scan() {
		line := hwa.Explorer.Peek()

		// Ignore the comments
		if strings.HasPrefix(line, "#") {
			hwa.Explorer.Consume()
			continue
		}

		// Handle the Classes
		if strings.HasPrefix(line, "C ") {
			err := readClassSection(hwa.Explorer, hwa)
			if err != nil {
				return err
			}

			continue
		}

		// Handle the vendors
		if lineStartsWithHex(line) {
			// Its a vendor
			readVendorSection(hwa.Explorer, hwa)
			// hwa.Explorer.Consume()
			continue
		}

		hwa.Explorer.Consume()
	}

	return nil
}
