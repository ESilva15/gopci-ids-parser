package hwarchiver

import (
	"reflect"
	"testing"
)

func TestParseClassLine(t *testing.T) {
	input := "C 03 Some Cool Hardware Stuff"
	expected := &Class{
		ID:         3,
		Name:       "Some Cool Hardware Stuff",
		Subclasses: make(map[int64]*Subclass),
	}

	result, err := parseClassLine(input)
	if err != nil {
		t.Errorf("Got error: %+v", err)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\nExpected:\n%+v\nGot:\n%+v\n", expected, result)
	}
}

// Test parseClassLine with some bad input
func TestParseClassLineWithBadInput(t *testing.T) {
	input := "C 03Some Cool Hardware Stuff"

	_, err := parseClassLine(input)
	if err == nil {
		t.Errorf("Got error: %+v", err)
	}
}
