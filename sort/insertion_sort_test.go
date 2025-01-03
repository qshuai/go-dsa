package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
	"github.com/qshuai/go-dsa/utils"
)

func Test_InsertionSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			arr := utils.Copy(tt.Args)
			if InsertionSort(arr); !reflect.DeepEqual(arr, tt.Expected) {
				t.Errorf("InsertionSort() = %v, want: %v", arr, tt.Expected)
			}
		})
	}
}
