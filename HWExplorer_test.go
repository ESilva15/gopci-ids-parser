package hwarchiver

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestExplorer(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Errorf("Failed to get current test file path")
	}

	testDir := filepath.Dir(filename)
	testdataFile := filepath.Join(testDir, "testdata", "hwexplorer_data.txt")

	file, err := os.Open(testdataFile)
	if err != nil {
		t.Errorf("Failed to open file: %+v", err)
	}
	scanner := bufio.NewScanner(file)
	_ = newHWExplorer(scanner)
}
