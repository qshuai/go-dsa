package leetcode

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case-1",
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			name: "case-2",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			name: "case-3",
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
		{
			name: "case-4",
			nums: []int{1},
			want: 2,
		},
		{
			name: "case-5",
			nums: []int{2, 1},
			want: 3,
		},
		{
			name: "case-6",
			nums: []int{2, 2, 2, 2, 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
