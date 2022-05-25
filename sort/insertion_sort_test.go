package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
)

func Test_InsertionSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			InsertionSort(tt.Args)
			if !reflect.DeepEqual(tt.Expected, tt.Args) {
				t.Errorf("InsertionSort[%s] expected: %v, bug got: %v", tt.Name, tt.Expected, tt.Args)
			}
		})
	}
}
