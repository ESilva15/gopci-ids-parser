package hwarchiver_test

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"gopkg.in/yaml.v2"

	hwarchive "github.com/ESilva15/gopci-ids-parser"
)

func getTestDataDir(file string) string {
	_, filename, _, _ := runtime.Caller(0)

	testDir := filepath.Dir(filename)
	testdataFile := filepath.Join(testDir, "testdata", file)

	return testdataFile
}

func getExpectedYAML(t *testing.T) *hwarchive.HWArchive {
	expectedYAMLfilepath := getTestDataDir("good_expected_yaml_output_of_pci_ids.yaml")

	expectedHWArchive := hwarchive.CreateHWArchive()
	yamlFile, err := os.ReadFile(expectedYAMLfilepath)
	if err != nil {
		t.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, expectedHWArchive)
	if err != nil {
		t.Fatal(err)
	}

	return expectedHWArchive
}

func TestHWArchiver(t *testing.T) {
	filePath := getTestDataDir("pci.ids")

	archive := hwarchive.CreateHWArchive()
	err := archive.Load(filePath)
	if err != nil {
		t.Error(err)
	}

	expectedArchive := getExpectedYAML(t)

	if !reflect.DeepEqual(expectedArchive, archive) {
		t.Errorf("Expected output doesn't match the result")
	}
}

func BenchmarkWithSetup(b *testing.B) {
	filePath := getTestDataDir("pci.ids")

	b.ResetTimer()
	for b.Loop() {
		archive := hwarchive.CreateHWArchive()
		archive.Load(filePath)
	}
}
