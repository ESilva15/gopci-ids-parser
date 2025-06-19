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
// ...<hex> <string>
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

func lineStartsWithHex(s string) bool {
	hexOffset, err := findHexOffset(s)
	if err != nil {
		return false
	}

	_, err = strconv.ParseInt(s[:hexOffset], 16, 64)
	if err != nil {
		return false
	}

	return true
}

func readTopLevelSection[
	P any,        // Top-level type: *Class, *Vendor
	C any,        // Nested type: *Subclass, *Device
](
	exp *HWExplorer,
	hwa *HWArchive,
	parseTop func(string) (P, error),
	addTop func(*HWArchive, P),
	nested func(P) error,
	printTop func(P),
) error {
	// Parse and add top-level object
	newTop, err := parseTop(exp.Peek())
	if err != nil {
		return err
	}

	printTop(newTop)
	addTop(hwa, newTop)
	exp.Consume()

	for exp.Scan() {
		line := exp.Peek()

		if strings.HasPrefix(line, "C ") || lineStartsWithHex(line) {
			break
		}
		if strings.HasPrefix(line, "\t\t") {
			return fmt.Errorf("line `%d` went from top to sub-sub-level", exp.line)
		}
		if strings.HasPrefix(line, "\t") {
			if err := nested(newTop); err != nil {
				return err
			}
			continue
		}

		exp.Consume()
	}

	return nil
}

func readSection[P any, C any](
	exp *HWExplorer,
	parent P,
	parse func(string) (C, error),
	add func(P, C) error,
	abort func(string) bool,
	match func(string) bool,
	parseNested func(C, string) error,
	print func(C),
	printNested func(C),
) error {
	newChild, err := parse(strings.TrimSpace(exp.Peek()))
	if err != nil {
		return err
	}

	print(newChild)
	if err := add(parent, newChild); err != nil {
		return err
	}

	exp.Consume()

	for exp.Scan() {
		line := exp.Peek()

		if abort(line) {
			break
		}

		if match(line) {
			if err := parseNested(newChild, strings.TrimSpace(line)); err != nil {
				return err
			}
			printNested(newChild)
			exp.Consume()
			continue
		}

		exp.Consume()
	}

	return nil
}
