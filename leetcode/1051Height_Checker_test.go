package leetcode

import (
	"testing"
)

func Test_heightChecker(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "case-1",
			args: []int{1, 1, 4, 2, 1, 3},
			want: 3,
		},
		{
			name: "case-2",
			args: []int{5, 1, 2, 3, 4},
			want: 5,
		},
		{
			name: "case-3",
			args: []int{1, 2, 3, 4, 5},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heightChecker(tt.args); got != tt.want {
				t.Errorf("heightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
