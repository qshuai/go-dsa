package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
	"github.com/qshuai/go-dsa/utils"
)

func TestQuickSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			arr := utils.Copy(tt.Args)
			QuickSort(arr, PartitionHead)
			if !reflect.DeepEqual(arr, tt.Expected) {
				t.Errorf("QuickSort-PartitionHead() = %v, want: %v", arr, tt.Expected)
			}

			arr = utils.Copy(tt.Args)
			QuickSort(arr, PartitionMiddle)
			if !reflect.DeepEqual(arr, tt.Expected) {
				t.Errorf("QuickSort-PartitionMiddle() = %v, want: %v", arr, tt.Expected)
			}

			arr = utils.Copy(tt.Args)
			QuickSort(arr, PartitionTail)
			if !reflect.DeepEqual(arr, tt.Expected) {
				t.Errorf("QuickSort-PartitionTail() = %v, want: %v", arr, tt.Expected)
			}

			arr = utils.Copy(tt.Args)
			QuickSort(arr, PartitionMedian)
			if !reflect.DeepEqual(arr, tt.Expected) {
				t.Errorf("QuickSort-PartitionMedian() = %v, want: %v", arr, tt.Expected)
			}
		})
	}
}
