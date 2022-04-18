package leetcode

import (
	"testing"
)

func Test_maximumGap(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "case-1",
			args: []int{3, 6, 9, 1},
			want: 3,
		},
		{
			name: "case-2",
			args: []int{10},
			want: 0,
		},
		{
			name: "case-3",
			args: []int{1, 1, 1, 1},
			want: 0,
		},
		{
			name: "case-4",
			args: []int{100, 3, 2, 1},
			want: 97,
		},
		{
			name: "case-5",
			args: []int{1, 1, 1, 1, 1, 5, 5, 5, 5, 5},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGap(tt.args); got != tt.want {
				t.Errorf("maximumGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
