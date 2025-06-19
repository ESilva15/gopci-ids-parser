package hwarchiver

import (
	"fmt"
	"testing"
)

// Test validHexChar with lower case inputs
func TestValidHexCharWithCorrectLowerCaseInputs(t *testing.T) {
	for ch := 'a'; ch <= 'f'; ch++ {
		result := validHexChar(ch)
		if !result {
			t.Errorf("Invalid output `%+v` for input `%+v` (%c)", result, ch, ch)
		}
	}
}

// Test validHexChar with upper case inputs
func TestValidHexCharWithCorrectUpperCaseInputs(t *testing.T) {
	for ch := 'A'; ch <= 'F'; ch++ {
		result := validHexChar(ch)
		if !result {
			t.Errorf("Invalid output `%+v` for input `%+v` (%c)", result, ch, ch)
		}
	}
}

// Test validHexChar with digits
func TestValidHexCharWithDigits(t *testing.T) {
	for ch := '0'; ch <= '9'; ch++ {
		result := validHexChar(ch)
		if !result {
			t.Errorf("Invalid output `%+v` for input `%+v` (%c)", result, ch, ch)
		}
	}
}

// Test validHexChar with wrong lower case chars
func TestValidHexCharWithBadLowerCaseChars(t *testing.T) {
	for ch := 'g'; ch <= 'z'; ch++ {
		result := validHexChar(ch)
		if result {
			t.Errorf("Invalid output `%+v` for input `%+v` (%c)", result, ch, ch)
		}
	}
}

// Test validHexChar with wrong upper case chars
func TestValidHexCharWithBadUpperCaseChars(t *testing.T) {
	for ch := 'G'; ch <= 'Z'; ch++ {
		result := validHexChar(ch)
		if result {
			t.Errorf("Invalid output `%+v` for input `%+v` (%c)", result, ch, ch)
		}
	}
}

// Test findHexOffset
func TestFindHexOffset(t *testing.T) {
	failCount := 0
	for k := range 0xFF {
		input := fmt.Sprintf("%02x ", k)
		expectedOffset := 2
		if k > 0xFF {
			expectedOffset = 3
		}
		if k > 0xFFF {
			expectedOffset = 4
		}

		result, err := findHexOffset(input)
		if err != nil {
			t.Errorf("Got error: %+v", err)
			failCount++
		}

		if result != expectedOffset {
			t.Errorf("Expected: %+v [%s], Got: %+v [%s]", expectedOffset, input,
				result, input[:result])
			failCount++
		}

		if failCount > 10 {
			break
		}
	}
}

// Test findHexOffset with some bad input
func TestFindHexOffsetBadInput(t *testing.T) {
	input := "zA"

	_, err := findHexOffset(input)
	if err == nil {
		t.Errorf("Got error: %+v", err)
	}
}

// test parseHexFieldsLineWithBadHex with good expected input
func TestParseHexFieldsLine(t *testing.T) {
	input := "0003 0005 A nice line"
	expectHex1 := int64(3)
	expectHex2 := int64(5)
	expectName := "A nice line"

	var hex1 int64
	var hex2 int64
	name, err := parseHexFieldsLine(input, &hex1, &hex2)
	if err != nil {
		t.Errorf("Got an error: %+v", err)
	}

	if hex1 != expectHex1 {
		t.Errorf("Expected: %+v, Got: %+v", expectHex1, hex1)
	}
	if hex2 != expectHex2 {
		t.Errorf("Expected: %+v, Got: %+v", expectHex2, hex2)
	}
	if name != expectName {
		t.Errorf("Expected: %+v, Got: %+v", expectName, name)
	}
}

// test parseHexFieldsLineWithBadHex with bad hex numbers
func TestParseHexFieldsLineWithBadHex(t *testing.T) {
	input := "00Z3 0005 A nice line"

	var hex1 int64
	var hex2 int64
	_, err := parseHexFieldsLine(input, &hex1, &hex2)
	if err == nil {
		t.Errorf("Got an error: %+v", err)
	}
}

// test parseHexFieldsLineWithBadHex with fewer hexes than requested
func TestParseHexFieldsLineWithBadHexCount(t *testing.T) {
	input := "00Z3 0005"

	var hex1 int64
	var hex2 int64
	_, err := parseHexFieldsLine(input, &hex1, &hex2)
	if err == nil {
		t.Errorf("Got an error: %+v", err)
	}
}

// Test lineStartsWithHex with good input
func TestLineStartsWithHex(t *testing.T) {
	input := "00AB A Cool Line"
	expected := true

	result := lineStartsWithHex(input)

	if result != expected {
		t.Errorf("Expected: %+v, Got: %+v", expected, result)
	}
}

// Test lineStartsWithHex without a hex
func TestLineStartsWithHexWithoutHex(t *testing.T) {
	input := "Z Cool Line"
	expected := false

	result := lineStartsWithHex(input)

	if result != expected {
		t.Errorf("Expected: %+v, Got: %+v", expected, result)
	}
}

// Test lineStartsWithHex with bad hex
func TestLineStartsWithHexWithBadHex(t *testing.T) {
	input := "12G4 Cool Line"
	expected := false

	result := lineStartsWithHex(input)

	if result != expected {
		t.Errorf("Expected: %+v, Got: %+v", expected, result)
	}
}
