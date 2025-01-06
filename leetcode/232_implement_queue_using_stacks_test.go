package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func TestMyQueue(t *testing.T) {
	q := NewMyQueue[int]()
	q.Push(1)
	q.Push(2)
	utils.AssertEqual(t, q.Peek(), 1)
	utils.AssertEqual(t, q.Pop(), 1)
	utils.AssertFalse(t, q.Empty())

	q = NewMyQueue[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	utils.AssertEqual(t, q.Pop(), 1)
	q.Push(5)
	utils.AssertEqual(t, q.Pop(), 2)
	utils.AssertEqual(t, q.Pop(), 3)
	utils.AssertEqual(t, q.Pop(), 4)
	utils.AssertEqual(t, q.Pop(), 5)
}
