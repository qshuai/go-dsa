package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		sorted []int
	}{
		{
			name:   "standard example",
			array:  []int{50, 10, 90, 30, 70, 40, 80, 60, 20},
			sorted: []int{10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
	}

	for _, test := range tests {
		ret := QuickSort(test.array)
		if !reflect.DeepEqual(ret, test.sorted) {
			t.Errorf("%s, want: %v, but got: %v",
				test.name, test.sorted, ret)
		}
	}
}
