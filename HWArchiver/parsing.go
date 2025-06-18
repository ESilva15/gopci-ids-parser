package hwarchiver

import (
	"fmt"
	"strconv"
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

// parseHexStringLine parses a line with the following syntax
// '<hex> <string>'
func parseHexStringLine(hex *int64, name *string, input string) error {
	hexOffset, err := findHexOffset(input)
	if err != nil {
		return err
	}

	*hex, err = strconv.ParseInt(input[:hexOffset], 16, 64)
	if err != nil {
		return err
	}
	*name = input[hexOffset+1:]

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
