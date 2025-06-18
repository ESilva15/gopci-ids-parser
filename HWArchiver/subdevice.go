package hwarchiver

import (
	"strconv"
	"strings"
)

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
	hexOffset, err := findHexOffset(s)
	if err != nil {
		return nil, err
	}
	firstHex, err := strconv.ParseInt(s[:hexOffset], 16, 64)
	if err != nil {
		return nil, err
	}

	substr := strings.TrimSpace(s[hexOffset:])
	hexOffset, err = findHexOffset(substr)
	if err != nil {
		return nil, err
	}

	secondHex, err := strconv.ParseInt(strings.TrimSpace(substr[:hexOffset]), 16, 64)
	if err != nil {
		return nil, err
	}

	name := strings.TrimSpace(substr[secondHex:])

	return &Subdevice{firstHex, secondHex, name}, nil
}
