package types

type Set[T comparable] struct {
	container map[T]struct{}
}

func (s *Set[T]) Add(ele T) {
	s.container[ele] = struct{}{}
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
