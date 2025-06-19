package hwarchiver

import (
	"testing"
)

func TestAddToMap_Success(t *testing.T) {
	m := make(map[int64]string)
	err := AddToMap(m, int64(1), "value", "TestPath")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if v, ok := m[1]; !ok || v != "value" {
		t.Fatalf("expected map[1] = 'value', got %v", m[1])
	}
}

func TestAddToMap_AlreadyExists(t *testing.T) {
	m := map[int64]string{1: "existing"}
	err := AddToMap(m, int64(1), "new", "TestPath")

	if err == nil {
		t.Fatal("expected error for duplicate key, got nil")
	}
}

func TestAddToMap_NilMap(t *testing.T) {
	var m map[int64]string
	err := AddToMap(m, int64(1), "value", "TestPath")

	if err == nil {
		t.Fatal("expected error for nil map, got nil")
	}
}
