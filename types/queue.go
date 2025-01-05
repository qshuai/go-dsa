package types

type Queue[T any] struct {
	container []T
}

func (q *Queue[T]) Push(eles ...T) {
	q.container = append(q.container, eles...)
}

func (q *Queue[T]) Pop(eles ...T) (v T, ok bool) {
	if len(q.container) == 0 {
		return
	}

	ok = true
	v = q.container[0]
	q.container = q.container[1:]
	return
}

func (q *Queue[T]) Peek(eles ...T) (v T, ok bool) {
	if len(q.container) == 0 {
		return
	}

	ok = true
	v = q.container[0]
	return
}

func (q *Queue[T]) Size(eles ...T) int {
	return len(q.container)
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func NewQueueWithCapacity[T any](capacity int) *Queue[T] {
	return &Queue[T]{container: make([]T, 0, capacity)}
}
