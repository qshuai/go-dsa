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
				t.Errorf("QuickSort-PartitionHead[%s] Expected: %v but got: %v", tt.Name, tt.Expected, tt.Args)
			}

			QuickSort(tt.Args, PartitionMiddle)
			if !reflect.DeepEqual(tt.Args, tt.Expected) {
				t.Errorf("QuickSort-PartitionMiddle[%s] Expected: %v but got: %v", tt.Name, tt.Expected, tt.Args)
			}

			QuickSort(tt.Args, PartitionTail)
			if !reflect.DeepEqual(tt.Args, tt.Expected) {
				t.Errorf("QuickSort-PartitionTail[%s] Expected: %v but got: %v", tt.Name, tt.Expected, tt.Args)
			}

			QuickSort(tt.Args, PartitionMedian)
			if !reflect.DeepEqual(tt.Args, tt.Expected) {
				t.Errorf("QuickSort-PartitionMedian[%s] Expected: %v but got: %v", tt.Name, tt.Expected, tt.Args)
			}
		})
	}
}
