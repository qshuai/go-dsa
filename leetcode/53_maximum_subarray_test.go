package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "case-1",
			args: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "case-2",
			args: []int{1},
			want: 1,
		},
		{
			name: "case-3",
			args: []int{5, 4, -1, 7, 8},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
