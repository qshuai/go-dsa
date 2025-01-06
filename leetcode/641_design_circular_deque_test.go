package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func TestCircularDeque(t *testing.T) {
	cd := Constructor(3)
	utils.AssertTrue(t, cd.InsertLast(1))
	utils.AssertTrue(t, cd.InsertLast(2))
	utils.AssertTrue(t, cd.InsertFront(3))
	utils.AssertFalse(t, cd.InsertFront(4))
	utils.AssertEqual(t, cd.GetRear(), 2)
	utils.AssertTrue(t, cd.IsFull())
	utils.AssertTrue(t, cd.DeleteLast())
	utils.AssertTrue(t, cd.InsertFront(4))
	utils.AssertEqual(t, cd.GetFront(), 4)
}

func TestCircularDeque2(t *testing.T) {
	cd := Constructor(2)
	utils.AssertTrue(t, cd.InsertFront(7))
	utils.AssertTrue(t, cd.DeleteLast())
	utils.AssertEqual(t, cd.GetFront(), -1)
	utils.AssertTrue(t, cd.InsertLast(5))
	utils.AssertTrue(t, cd.InsertFront(0))
	utils.AssertEqual(t, cd.GetFront(), 0)
	utils.AssertEqual(t, cd.GetRear(), 5)
	utils.AssertEqual(t, cd.GetFront(), 0)
	utils.AssertEqual(t, cd.GetFront(), 0)
	utils.AssertEqual(t, cd.GetRear(), 5)
	utils.AssertFalse(t, cd.InsertLast(0))
}
