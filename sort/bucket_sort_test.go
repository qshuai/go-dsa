package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
	"github.com/qshuai/go-dsa/utils"
)

func TestBucketSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			arr := utils.Copy(tt.Args)
			if got := BucketSort(arr); !reflect.DeepEqual(got, tt.Expected) {
				t.Errorf("BucketSort() = %v, want: %v", got, tt.Expected)
			}
		})
	}
}
