package gog

// KV represents a key-value pair.
type KV[K comparable, V any] struct {
	// Key holds the key of this key-value pair.
	Key K

	// Value holds the value of this key-value pair.
	Value V
}

// Entries returns a slice of key-value pair mappings that make up
// the specified map.
//
// Note: There is no guarantee as to the order of the returned entries, nor
// is that order guaranteed to be consistent across two calls.
func Entries[K comparable, V any](m map[K]V) []KV[K, V] {
	result := make([]KV[K, V], 0, len(m))
	for k, v := range m {
		result = append(result, KV[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return result
}
