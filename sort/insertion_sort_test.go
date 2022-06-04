package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
)

func Test_InsertionSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			if InsertionSort(tt.Args); !reflect.DeepEqual(tt.Expected, tt.Args) {
				t.Errorf("InsertionSort() = %v, want: %v", tt.Args, tt.Expected)
			}
		})
	}
}
