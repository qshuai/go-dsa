package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
)

func Test_mergeSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			if got := MergeSort(tt.Args); !reflect.DeepEqual(got, tt.Expected) {
				t.Errorf("MergeSort() = %v, want: %v", got, tt.Expected)
			}
		})
	}
}
