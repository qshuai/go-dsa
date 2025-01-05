package types

type Stack[T any] struct {
	container []T
}

func (s *Stack[T]) Push(eles ...T) {
	s.container = append(s.container, eles...)
}

func (s *Stack[T]) Pop() (v T, ok bool) {
	if len(s.container) == 0 {
		return
	}

	ok = true
	v = s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]
	return
}

func (s *Stack[T]) Peek() (v T, ok bool) {
	if len(s.container) == 0 {
		return
	}

	ok = true
	v = s.container[len(s.container)-1]
	return
}

func (s *Stack[T]) Size() int {
	return len(s.container)
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func NewStackWithCapacity[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		container: make([]T, 0, capacity),
	}
}
