package leetcode

import (
	"testing"
)

func Test_hIndex(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "case-1",
			args: []int{3, 0, 6, 1, 5},
			want: 3,
		},
		{
			name: "case-2",
			args: []int{1, 3, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
