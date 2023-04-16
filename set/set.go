package set

// A Set is alis that contains no duplicate elements
type Set[T comparable] struct {
	// List of Items of type T in the set
	Items map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		Items: make(map[T]bool),
	}
}

// Add adds an item to the set
func (s Set[T]) Add(data T) {

	if s.Items[data] {
		return // Already in the map
	}

	s.Items[data] = true
}

// Del Deletes an item from the set
func (s Set[T]) Del(data T) bool {

	if s.Items[data] {
		delete(s.Items, data)
		return true
	}

	return false
}
