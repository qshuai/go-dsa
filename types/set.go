package types

// Set is implemented in Go, using the Go standard library’s map
// Set is not concurrency safe
type Set[T comparable] struct {
	container map[T]struct{}
}

func (s *Set[T]) Add(eles ...T) {
	for _, ele := range eles {
		s.container[ele] = struct{}{}
	}
}

func (s *Set[T]) Remove(eles ...T) {
	for _, ele := range eles {
		delete(s.container, ele)
	}
}

func (s *Set[T]) Has(ele T) bool {
	_, ok := s.container[ele]
	return ok
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		container: make(map[T]struct{}),
	}
}

func NewSetWithCapacity[T comparable](capacity int) *Set[T] {
	return &Set[T]{
		container: make(map[T]struct{}, capacity),
	}
}
