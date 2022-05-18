package sort

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "case-1",
			args: []int{2, 8, 3, 5},
			want: []int{2, 3, 5, 8},
		},
		{
			name: "case-2",
			args: []int{5, 1, 2, 1, 1, 1, 1, 1, 1, 1, 3},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BucketSort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
