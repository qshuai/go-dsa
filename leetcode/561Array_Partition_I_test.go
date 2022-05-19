package leetcode

import (
	"testing"
)

func Test_arrayPairSum(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "case-1",
			args: []int{1, 4, 3, 2},
			want: 4,
		},
		{
			name: "case-1",
			args: []int{6, 2, 6, 5, 1, 2},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSum(tt.args); got != tt.want {
				t.Errorf("arrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
