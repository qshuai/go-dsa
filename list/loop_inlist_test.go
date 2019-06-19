package list

import (
	"testing"
)

func TestContainLoopInList(t *testing.T) {
	// only one node
	l1 := NewNodeObject(1, nil)
	if containLoopInList(l1) {
		t.Errorf("test case: only on node, not loop but got have loop")
	}

	// two node list
	l2 := NewNodeObject(1, NewNodeObject(2, nil))
	if containLoopInList(l2) {
		t.Errorf("test case: two nodes list, not loop but got have loop")
	}

	// many node list without loop
	l3 := NewNodeObject(1, NewNodeObject(2, NewNodeObject(3, NewNodeObject(4, nil))))
	if containLoopInList(l3) {
		t.Errorf("test case: arbitrarily nodes list, not loop but got have loop")
	}

	// list with loop
	l4 := NewNodeObject(1, NewNodeObject(2, NewNodeObject(3, NewNodeObject(4, nil))))
	l4.Next.Next.Next.Next = l4
	if !containLoopInList(l4) {
		t.Errorf("test case: list with loop, have loop get no loop")
	}
}
