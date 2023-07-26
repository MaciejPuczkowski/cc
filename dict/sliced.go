package dict

// SlicedWrapper[K comparable, V any] is a builder that helps working with dicts containing slices.
// It provides Append method that appends an element to a slice under the given key.

type SlicedWrapper[K comparable, V any] struct {
	d GetFSetter[K, []V]
}

// Sliced creates a new SlicedWrapper.
func Sliced[K comparable, V any](d GetFSetter[K, []V]) *SlicedWrapper[K, V] {
	return &SlicedWrapper[K, V]{d}
}

// Append appends an element to a slice under the given key.
// If the slice does not exist, it is created.
func (s *SlicedWrapper[K, V]) Append(k K, v V) {
	s.d.Set(k, append(s.d.GetF(k, func() []V { return []V{} }), v))
}
