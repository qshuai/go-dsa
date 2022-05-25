package search

import (
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func TestContainLoopInList(t *testing.T) {
	// only one node
	l1 := types.NewListNode(1, 1)
	if ContainLoopInList(l1) {
		t.Errorf("test case: only on node, not loop but got have loop")
	}

	// two node list
	l2 := types.NewListNode(1, 2)
	if ContainLoopInList(l2) {
		t.Errorf("test case: two nodes list, not loop but got have loop")
	}

	// many node list without loop
	l3 := types.NewListNode(1, 4)
	if ContainLoopInList(l3) {
		t.Errorf("test case: arbitrarily nodes list, not loop but got have loop")
	}

	// list with loop
	l4 := types.NewListNode(1, 4)
	l4.Next.Next.Next = l4.Next
	if !ContainLoopInList(l4) {
		t.Errorf("test case: list with loop, have loop get no loop")
	}
}
