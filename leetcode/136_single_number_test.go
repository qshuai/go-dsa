package leetcode

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case-1",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "case-2",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "case-3",
			nums: []int{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
