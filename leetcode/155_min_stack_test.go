package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func TestMinStack(t *testing.T) {
	ms := NewMinStack[int]()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	utils.AssertEqual(t, ms.GetMin(), -3)

	ms.Pop()
	utils.AssertEqual(t, ms.Top(), 0)
	utils.AssertEqual(t, ms.GetMin(), -2)
}
