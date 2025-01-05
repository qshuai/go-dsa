package types

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func TestQueue_PushPop(t *testing.T) {
	queue := NewQueue[int]()
	utils.AssertEqual(t, queue.Size(), 0)

	queue.Push(1)
	utils.AssertEqual(t, queue.Size(), 1)

	queue.Push(2)
	utils.AssertEqual(t, queue.Size(), 2)

	queue.Push(3, 4)
	utils.AssertEqual(t, queue.Size(), 4)

	v, ok := queue.Pop()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, 1)
	utils.AssertEqual(t, queue.Size(), 3)

	v, ok = queue.Pop()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, 2)
	utils.AssertEqual(t, queue.Size(), 2)

	v, ok = queue.Pop()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, 3)
	utils.AssertEqual(t, queue.Size(), 1)

	v, ok = queue.Pop()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, 4)
	utils.AssertEqual(t, queue.Size(), 0)

	v, ok = queue.Pop()
	utils.AssertFalse(t, ok)
	utils.AssertEqual(t, v, 0)
	utils.AssertEqual(t, queue.Size(), 0)
}

func TestQueue_Peek(t *testing.T) {
	queue := NewQueueWithCapacity[int](5)
	utils.AssertEqual(t, queue.Size(), 0)

	v, ok := queue.Peek()
	utils.AssertFalse(t, ok)
	utils.AssertEqual(t, v, 0)
	utils.AssertEqual(t, queue.Size(), 0)

	queue.Push(1, 2, 3)
	utils.AssertEqual(t, queue.Size(), 3)

	v, ok = queue.Peek()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, 1)
	utils.AssertEqual(t, queue.Size(), 3)
}
