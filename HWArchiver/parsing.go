package hwarchiver

import (
	"fmt"
	"strconv"
	"strings"
)

func validHexChar(c rune) bool {
	return (c >= 'a' && c <= 'f') ||
		(c >= 'A' && c <= 'F') ||
		(c >= '0' && c <= '9')
}

func findHexOffset(s string) (int, error) {
	offset := 0
	for i, c := range s {
		if c == ' ' {
			offset = i
			break
		}

		if !validHexChar(c) {
			return -1, fmt.Errorf("invalid hex char `%c`", c)
		}
	}

	return offset, nil
}

// parseHexFieldsLine will separate the string by its spaces and then
// parse N hex at the beginning of the line and join the remaining string
func parseHexFieldsLine(s string, hexOut ...*int64) (string, error) {
	fields := strings.Fields(s)
	if len(fields) < len(hexOut)+1 {
		return "", fmt.Errorf("expected at least %d fields, got %d",
			len(hexOut)+1, len(fields))
	}

	for i, ptr := range hexOut {
		val, err := strconv.ParseInt(fields[i], 16, 64)
		if err != nil {
			return "", fmt.Errorf("invalid hex at position %d: %v", i, err)
		}
		*ptr = val
	}

	// Join the remaining fields as the name
	name := strings.Join(fields[len(hexOut):], " ")
	return name, nil
}

// parseHexHexStringLine parses a line with the following syntax
// '<hex> <hex> <string>'
func parseHexHexStringLine(hex1 *int64, hex2 *int64, name *string, s string) error {
	retName, err := parseHexFieldsLine(s, hex1, hex2)
	if err != nil {
		return err
	}
	*name = retName

	return nil
}

func lineStartsWithHex(s string) bool {
	// This code is repeated with parseHexStringLine - check what we can do
	hexOffset, err := findHexOffset(s)
	if err != nil {
		return false
	}

	_, err = strconv.ParseInt(s[:hexOffset], 16, 64)
	if err != nil {
		return false
	}
	// to here

	return true
}
