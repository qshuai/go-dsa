package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func TestMyStack(t *testing.T) {
	stack := NewMyStack[int]()
	stack.Push(1)
	stack.Push(2)
	utils.AssertEqual(t, stack.Top(), 2)
	utils.AssertEqual(t, stack.Pop(), 2)
	utils.AssertFalse(t, stack.Empty())

	stack = NewMyStack[int]()
	stack.Push(1)
	stack.Push(2)
	utils.AssertEqual(t, stack.Pop(), 2)
	stack.Push(3)
	stack.Push(4)
	utils.AssertEqual(t, stack.Pop(), 4)
	utils.AssertEqual(t, stack.Pop(), 3)
}
