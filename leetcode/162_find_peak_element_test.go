package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func Test_findPeakElement(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "case-1",
			args: []int{1, 2, 3, 1},
			want: []int{2},
		},
		{
			name: "case-2",
			args: []int{1, 2, 1, 3, 5, 6, 4},
			want: []int{1, 5},
		},
		{
			name: "case-3",
			args: []int{3, 4, 3, 2, 1},
			want: []int{1},
		},
		{
			name: "case-4",
			args: []int{3, 4, 5, 6, 3},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args); !utils.Contain(tt.want, got) {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
