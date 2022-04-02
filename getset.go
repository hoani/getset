package getset

type Set[T comparable] map[T]struct{}

// New creates a new set with the given items.
// Items must all be the same type.
func New[T comparable](items ...T) Set[T] {
	m := make(map[T]struct{}, len(items))
	for _, item := range items {
		m[item] = struct{}{}
	}
	return Set[T](m)
}

// Has checks whether an item is in the given set.
func (s Set[T]) Has(item T) bool {
	_, has := s[item]
	return has
}

// Insert adds another value to the set. If the value already exists, the set remains the same.
func (s Set[T]) Insert(item T) {
	s[item] = struct{}{}
}

// ToArray returns a slice of all items in the set.
func (s Set[T]) ToArray() []T {
	result := make([]T, 0, len(s))
	for item := range s {
		result = append(result, item)
	}
	return result
}
