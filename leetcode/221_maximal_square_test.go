package leetcode

import (
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	tests := []struct {
		name string
		args [][]byte
		want int
	}{
		{
			name: "case-1",
			args: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'}},
			want: 4,
		},
		{
			name: "case-2",
			args: [][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			want: 1,
		},
		{
			name: "case-3",
			args: [][]byte{{'0'}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
