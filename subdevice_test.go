package hwarchiver

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseSubdeviceLine(t *testing.T) {
	input := "		001c 0005  2 Channel CAN Bus SJC1000 (Optically Isolated)"
	expected := &Subdevice{
		Identity: Identity{
			ID:   0x001c,
			Name: "2 Channel CAN Bus SJC1000 (Optically Isolated)",
		},
		Subdevice: 0x0005,
	}

	result, err := parseSubdeviceLine(strings.TrimSpace(input))
	if err != nil {
		t.Errorf("Got an error: %+v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nExpected:\n%+v\nGot:\n%+v\n", expected, result)
	}
}

func TestParseSubdeviceLineWithBadInput(t *testing.T) {
	input := "		001G 0005  2 Channel CAN Bus SJC1000 (Optically Isolated)"

	_, err := parseSubdeviceLine(strings.TrimSpace(input))
	if err == nil {
		t.Errorf("Got an error: %+v", err)
	}
}
