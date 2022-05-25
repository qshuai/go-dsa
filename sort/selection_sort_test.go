package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
)

func Test_SelectionSort(t *testing.T) {
	for _, test := range testdata.GetTestCases() {
		SelectionSort(test.Args)
		if !reflect.DeepEqual(test.Args, test.Expected) {
			t.Errorf("SelectionSort[%s] expected: %v, but got: %v", test.Name, test.Expected, test.Args)
		}
	}
}
