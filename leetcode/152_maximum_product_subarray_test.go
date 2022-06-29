package leetcode

import "testing"

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "case-1",
			args: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "case-2",
			args: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "case-3",
			args: []int{0, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
