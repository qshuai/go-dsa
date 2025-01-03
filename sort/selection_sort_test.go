package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
	"github.com/qshuai/go-dsa/utils"
)

func Test_SelectionSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			arr := utils.Copy(tt.Args)
			if SelectionSort(arr); !reflect.DeepEqual(arr, tt.Expected) {
				t.Errorf("SelectionSort() = %v, want: %v", arr, tt.Expected)
			}
		})
	}
}
