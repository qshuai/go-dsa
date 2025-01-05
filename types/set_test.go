package types

import "testing"

func TestSet(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)

	if !set.Has(1) {
		t.Fatalf("set should have ele: 1")
	}
	if set.Has(2) {
		t.Fatalf("set should not have ele: 2")
	}

	set.Remove(1)
	if set.Has(1) {
		t.Fatalf("element should have been removed: 1")
	}
}

func TestSetWithCapacity(t *testing.T) {
	set := NewSetWithCapacity[int](3)
	set.Add(1, 2, 3)

	if !set.Has(1) {
		t.Fatalf("set should have ele: 1")
	}
	if !set.Has(2) {
		t.Fatalf("set should have ele: 2")
	}
	if !set.Has(3) {
		t.Fatalf("set should have ele: 3")
	}
	if set.Has(4) {
		t.Fatalf("set should not have ele: 4")
	}

	set.Add(4)
	if !set.Has(4) {
		t.Fatalf("set should have ele: 4")
	}
	if !set.Has(1) {
		t.Fatalf("set should have ele: 1")
	}

	set.Remove(2)
	if set.Has(2) {
		t.Fatalf("element should have been removed: 2")
	}
}
