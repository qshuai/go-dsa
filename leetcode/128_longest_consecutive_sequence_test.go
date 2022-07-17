package leetcode

import (
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "case-1",
			args: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "case-2",
			args: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
