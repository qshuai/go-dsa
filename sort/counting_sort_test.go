package sort

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/sort/testdata"
)

func TestCountingSort(t *testing.T) {
	for _, tt := range testdata.GetTestCases() {
		t.Run(tt.Name, func(t *testing.T) {
			CountingSort(tt.Args, tt.Min, tt.Max)
			if !reflect.DeepEqual(tt.Expected, tt.Args) {
				t.Errorf("CountingSort[%s] expected: %v, bug got: %v", tt.Name, tt.Expected, tt.Args)
			}
		})
	}
}
