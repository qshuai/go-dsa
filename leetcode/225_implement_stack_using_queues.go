package leetcode

import "github.com/qshuai/go-dsa/types"

type MyStack[T any] struct {
	queue *types.Queue[T]
}

func NewMyStack[T any]() *MyStack[T] {
	return &MyStack[T]{
		queue: types.NewQueue[T](),
	}
}

func (ms *MyStack[T]) Push(x T) {
	ms.queue.Push(x)

	for i := 0; i < ms.queue.Size()-1; i++ {
		v, _ := ms.queue.Pop()
		ms.queue.Push(v)
	}
}

func (ms *MyStack[T]) Pop() T {
	v, _ := ms.queue.Pop()
	return v
}

func (ms *MyStack[T]) Top() T {
	v, _ := ms.queue.Peek()
	return v
}

func (ms *MyStack[T]) Empty() bool {
	return ms.queue.Size() == 0
}
