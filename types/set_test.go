package types

import "testing"

func TestSet_AddHas(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)

	if !set.Has(1) {
		t.Fatalf("set should have ele: 1")
	}
	if set.Has(2) {
		t.Fatalf("set should not have ele: 2")
	}
}
