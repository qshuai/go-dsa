package sort

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "general case",
			args: []int{3, 4, 1, 5, 3, 6, 5},
			want: []int{1, 3, 3, 4, 5, 5, 6},
		},
		{
			name: "sorted array",
			args: []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "single element",
			args: []int{34},
			want: []int{34},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
