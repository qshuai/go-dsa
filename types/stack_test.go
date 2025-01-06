package types

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func TestStack_PushPop(t *testing.T) {
	stack := NewStack[int]()
	utils.AssertEqual(t, stack.Size(), 0)
	utils.AssertTrue(t, stack.Empty())

	// push data
	stack.Push(1, 2, 3)

	utils.AssertEqual(t, stack.Size(), 3)

	v, ok := stack.Pop()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, 3)
	utils.AssertEqual(t, stack.Size(), 2)

	// push again
	stack.Push(4)

	utils.AssertEqual(t, stack.Size(), 3)

	v, ok = stack.Pop()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, 4)
	utils.AssertEqual(t, stack.Size(), 2)

	v, ok = stack.Pop()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, 2)
	utils.AssertEqual(t, stack.Size(), 1)

	v, ok = stack.Pop()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, 1)
	utils.AssertEqual(t, stack.Size(), 0)
	utils.AssertTrue(t, stack.Empty())

	v, ok = stack.Pop()
	utils.AssertFalse(t, ok)
	utils.AssertEqual(t, v, 0)
	utils.AssertEqual(t, stack.Size(), 0)
	utils.AssertTrue(t, stack.Empty())
}

func TestStack_Peek(t *testing.T) {
	stack := NewStackWithCapacity[string](5)
	utils.AssertEqual(t, stack.Size(), 0)
	utils.AssertTrue(t, stack.Empty())

	v, ok := stack.Peek()
	utils.AssertFalse(t, ok)
	utils.AssertEqual(t, v, "")
	utils.AssertEqual(t, stack.Size(), 0)

	// push data
	stack.Push("hello", "world", "!")

	// peek
	v, ok = stack.Peek()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, "!")
	utils.AssertEqual(t, stack.Size(), 3)

	// peek again
	v, ok = stack.Peek()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, "!")
	utils.AssertEqual(t, stack.Size(), 3)

	// pop data
	v, ok = stack.Pop()
	utils.AssertTrue(t, ok)
	utils.AssertEqual(t, v, "!")
	utils.AssertEqual(t, stack.Size(), 2)
}
