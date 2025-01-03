package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
	"github.com/qshuai/go-dsa/utils"
)

func TestCountingSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			arr := utils.Copy(tt.Args)
			CountingSort(arr, tt.Min, tt.Max)
			if !reflect.DeepEqual(arr, tt.Expected) {
				t.Errorf("CountingSort() = %v, want: %v", arr, tt.Expected)
			}
		})
	}
}
