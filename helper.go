package hwarchiver

import "fmt"

func addToMap[K comparable, T any](m map[K]T, key K, value T, mapname string) error {
	if m == nil {
		return fmt.Errorf("map %s is nill when it should not be", mapname)
	}

	if _, exists := m[key]; exists {
		return fmt.Errorf("key %v is already present in %s", key, mapname)
	}

	m[key] = value

	return nil
}
