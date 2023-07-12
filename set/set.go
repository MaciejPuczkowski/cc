package set

// Set is a set implementation
type Set[T comparable] struct {
	elements map[T]int
}

// New creates a new empty set
func New[T comparable]() *Set[T] {
	return &Set[T]{
		elements: make(map[T]int),
	}
}

// From creates a new set from a list
func From[T comparable](list []T) *Set[T] {
	var elements = make(map[T]int)
	for _, item := range list {
		if _, ok := elements[item]; !ok {
			elements[item] = 1
		} else {
			elements[item]++
		}
	}
	return &Set[T]{
		elements: elements,
	}
}

// Add adds an item to the set
func (s *Set[T]) Add(item T) {
	s.elements[item] = 1
}

// Contains checks if an item is in the set
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.elements[item]
	return ok
}

// Remove removes an item from the set
func (s *Set[T]) Remove(item T) {
	delete(s.elements, item)
}

// List returns a list of all items in the set
func (s *Set[T]) List() []T {
	result := make([]T, len(s.elements))
	i := 0
	for item, value := range s.elements {
		if value > 0 {
			result[i] = item
			i++
		}
	}
	return result[:i]
}

// Size returns the size of the set
func (s *Set[T]) Size() int {
	return len(s.elements)
}

func Union[T comparable](others ...*Set[T]) *Set[T] {
	newSet := New[T]()
	for _, other := range others {
		for _, item := range other.List() {
			newSet.Add(item)
		}
	}

	return newSet
}

func AreEqual[T comparable](s1 *Set[T], s2 *Set[T]) bool {
	if s1.Size() != s2.Size() {
		return false
	}
	for _, item := range s1.List() {
		if !s2.Contains(item) {
			return false
		}
	}

	return true
}

func Product[T comparable](sets ...*Set[T]) *Set[T] {
	newSet := New[T]()
	m := map[T]int{}
	for _, set := range sets {
		for _, item := range set.List() {
			if _, ok := m[item]; !ok {
				m[item] = 1
			} else {
				m[item]++
			}
		}
	}
	for item, value := range m {
		if value == len(sets) {
			newSet.Add(item)
		}
	}
	return newSet
}

func Xor[T comparable](sets ...*Set[T]) *Set[T] {
	newSet := New[T]()
	m := map[T]int{}
	for _, set := range sets {
		for _, item := range set.List() {
			if _, ok := m[item]; !ok {
				m[item] = 1
			} else {
				m[item]++
			}
		}
	}
	for item, value := range m {
		if value < len(sets) {
			newSet.Add(item)
		}
	}
	return newSet
}

func Diff[T comparable](sets ...*Set[T]) *Set[T] {
	newSet := New[T]()
	m := map[T]int{}
	for _, set := range sets {
		for _, item := range set.List() {
			if _, ok := m[item]; !ok {
				m[item] = 1
			} else {
				m[item]++
			}
		}
	}
	for item, value := range m {
		if value == 1 {
			newSet.Add(item)
		}
	}
	return newSet
}
