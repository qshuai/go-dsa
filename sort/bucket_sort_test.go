package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
)

func TestBucketSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			if got := BucketSort(tt.Args); !reflect.DeepEqual(got, tt.Expected) {
				t.Errorf("BucketSort() = %v, want: %v", got, tt.Expected)
			}
		})
	}
}
