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

// Test parseHexStringLine with good inputs
func TestParseHexStringLine(t *testing.T) {
	input := "05 A Cool Name"
	expectedHex := int64(5)
	expectedName := "A Cool Name"

	var hex int64
	var name string
	err := parseHexStringLine(&hex, &name, input)
	if err != nil {
		t.Errorf("Got error: %+v", err)
	}

	if hex != expectedHex {
		t.Errorf("Expected: %+v, Got %+v", expectedHex, hex)
	}

	if name != expectedName {
		t.Errorf("Expected: %+v, Got %+v", expectedName, name)
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
