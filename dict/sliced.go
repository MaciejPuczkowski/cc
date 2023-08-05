package dict

import "encoding/json"

// SlicedWrapper[K comparable, V any] is a builder that helps working with dicts containing slices.
type SlicedWrapper[K comparable, V any] struct {
	d Dicter[K, []V]
}

// Sliced creates a new SlicedWrapper.
func Sliced[K comparable, V any](d Dicter[K, []V]) *SlicedWrapper[K, V] {
	return &SlicedWrapper[K, V]{d}
}

// Append appends an element to a slice under the given key.
// If the slice does not exist, it is created.
func (s *SlicedWrapper[K, V]) Append(k K, v V) {
	s.d.Set(k, append(s.d.GetF(k, func() []V { return []V{} }), v))
}

// Prepend prepends an element to a slice under the given key.
func (s *SlicedWrapper[K, V]) Prepend(k K, v V) {
	if s.d.Has(k) {
		s.d.Set(k, append([]V{v}, s.d.GetF(k, func() []V { return []V{} })...))
	} else {
		s.Append(k, v)
	}
}

func (s *SlicedWrapper[K, V]) First(k K, def V) V {
	l := s.d.GetF(k, func() []V { return []V{def} })
	if len(l) == 0 {
		return def
	}
	return l[0]
}

func (s *SlicedWrapper[K, V]) FirstF(k K, factory func() V) V {
	l := s.d.GetF(k, func() []V { return []V{factory()} })
	if len(l) == 0 {
		return factory()
	}
	return l[0]
}

func (s *SlicedWrapper[K, V]) Last(k K, def V) V {
	l := s.d.GetF(k, func() []V { return []V{def} })
	if len(l) == 0 {
		return def
	}
	return l[len(l)-1]
}

func (s *SlicedWrapper[K, V]) LastF(k K, factory func() V) V {
	l := s.d.GetF(k, func() []V { return []V{factory()} })
	if len(l) == 0 {
		return factory()
	}
	return l[len(l)-1]
}

func (s *SlicedWrapper[K, V]) LenOf(k K) int {
	return len(s.d.GetF(k, func() []V { return []V{} }))
}

func (s *SlicedWrapper[K, V]) RemoveFirst(k K) bool {
	l := s.d.GetF(k, func() []V { return []V{} })
	if len(l) == 0 {
		return false
	}
	s.d.Set(k, l[1:])
	return true
}

func (s *SlicedWrapper[K, V]) RemoveLast(k K) bool {
	l := s.d.GetF(k, func() []V { return []V{} })
	if len(l) == 0 {
		return false
	}
	s.d.Set(k, l[:len(l)-1])
	return true
}

func (s *SlicedWrapper[K, V]) ConsumeFirst(k K, v V) (V, bool) {
	obj := s.First(k, v)
	removed := s.RemoveFirst(k)
	return obj, removed
}

func (s *SlicedWrapper[K, V]) ConsumeLast(k K, v V) (V, bool) {
	obj := s.Last(k, v)
	removed := s.RemoveLast(k)
	return obj, removed
}

func (s *SlicedWrapper[K, V]) Get(k K, def []V) []V {
	return s.d.GetF(k, func() []V { return def })
}

func (s *SlicedWrapper[K, V]) GetF(k K, factory func() []V) []V {
	return s.d.GetF(k, factory)
}

func (s *SlicedWrapper[K, V]) Has(k K) bool {
	return s.d.Has(k)
}

func (s *SlicedWrapper[K, V]) Set(k K, v []V) {
	s.d.Set(k, v)
}
func (s *SlicedWrapper[K, V]) SetDefault(k K, def []V) {
	s.d.SetDefault(k, def)
}

func (s *SlicedWrapper[K, V]) SetDefaultF(k K, factory func() []V) {
	s.d.SetDefaultF(k, factory)
}

func (s *SlicedWrapper[K, V]) Delete(k K) {
	s.d.Delete(k)
}

func (s *SlicedWrapper[K, V]) Keys() []K {
	return s.d.Keys()
}

func (s *SlicedWrapper[K, V]) Values() [][]V {
	return s.d.Values()
}

func (s *SlicedWrapper[K, V]) Items() []Item[K, []V] {
	return s.d.Items()
}

func (s *SlicedWrapper[K, V]) Len() int {
	return s.d.Len()
}

func (s *SlicedWrapper[K, V]) DeleteItem(i Item[K, []V]) {
	s.d.DeleteItem(i)
}

func (s *SlicedWrapper[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.d)
}

func (s *SlicedWrapper[K, V]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.d)
}

func (s *SlicedWrapper[K, V]) SetItem(i Item[K, []V]) {
	s.d.SetItem(i)
}
