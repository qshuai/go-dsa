package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
)

func TestQuickSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			QuickSort(tt.Args, PartitionHead)
			if !reflect.DeepEqual(tt.Args, tt.Expected) {
				t.Errorf("QuickSort-PartitionHead() = %v, want: %v", tt.Args, tt.Expected)
			}

			QuickSort(tt.Args, PartitionMiddle)
			if !reflect.DeepEqual(tt.Args, tt.Expected) {
				t.Errorf("QuickSort-PartitionMiddle() = %v, want: %v", tt.Args, tt.Expected)
			}

			QuickSort(tt.Args, PartitionTail)
			if !reflect.DeepEqual(tt.Args, tt.Expected) {
				t.Errorf("QuickSort-PartitionTail() = %v, want: %v", tt.Args, tt.Expected)
			}

			QuickSort(tt.Args, PartitionMedian)
			if !reflect.DeepEqual(tt.Args, tt.Expected) {
				t.Errorf("QuickSort-PartitionMedian() = %v, want: %v", tt.Args, tt.Expected)
			}
		})
	}
}
