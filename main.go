package main

import (
	"fmt"
	"os"
	hwarchive "systemStatus/HWArchiver"
	// "path/filepath"
)

// func findHwmonPath() (string, error) {
// 	basePath := "/sys/class/drm/card0/device/hwmon"
// 	entries, err := os.ReadDir(basePath)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	for _, entry := range entries {
// 		if entry.IsDir() {
// 			return filepath.Join(basePath, entry.Name()), nil
// 		}
// 	}
//
// 	return "", fmt.Errorf("no hwmon directory found")
// }

func detectGPUS() {
	drmPath := "/sys/class/drm"

	entries, err := os.ReadDir(drmPath)
	if err != nil {
		return
		// return "", err
	}

	for _, entry := range entries {
		name := entry.Name()
		fmt.Println("entry:", name)
	}
}

func main() {
	hwarchive := hwarchive.CreateHWArchive()

	// hwarchive.Load("/usr/share/hwdata/pci.ids")
	err := hwarchive.Load("./pid.ids")
	if err != nil {
		panic(err)
	}

	fmt.Println("Vendors:")
	for id := range hwarchive.Vendors {
		curVendor := hwarchive.Vendors[id]
		fmt.Printf("0x%04x %s\n", curVendor.ID, curVendor.Name)
		for subId := range curVendor.Devices {
			curDevice := curVendor.Devices[subId]
			fmt.Printf("    0x%04x %s\n", curDevice.ID, curDevice.Name)
			for subDev := range curDevice.Subdevices {
				curSubDev := curDevice.Subdevices[subDev]
				fmt.Printf("        0x%04x 0x%04x %s\n", curSubDev.ID,
					curSubDev.Subdevice, curSubDev.Name)
			}
		}
	}

	fmt.Println("\nClasses:")
	for id := range hwarchive.Classes {
		curClass := hwarchive.Classes[id]
		fmt.Printf("0x%04x %s\n", curClass.ID, curClass.Name)
		for subId := range curClass.Subclasses {
			curSubClass := curClass.Subclasses[subId]
			fmt.Printf("    0x%04x %s\n", curSubClass.ID, curSubClass.Name)
			for progIf := range curSubClass.ProgInterfaces {
				curProgIf := curSubClass.ProgInterfaces[progIf]
				fmt.Printf("        0x%04x %s\n", curProgIf.ID, curProgIf.Name)
			}
		}
	}

	// detectGPUS()
	// hwmonPath, err := findHwmonPath()
	// if err != nil {
	// 	fmt.Println("Failed to find hwmon path:", err)
	// 	return
	// }
	//
	// fmt.Println("Hmonpath: ", hwmonPath)
}
