package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
)

func Test_SelectionSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			if SelectionSort(tt.Args); !reflect.DeepEqual(tt.Args, tt.Expected) {
				t.Errorf("SelectionSort() = %v, want: %v", tt.Args, tt.Expected)
			}
		})
	}
}
