package leetcode

import "testing"

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "case-1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "case-2",
			height: []int{1, 1},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Errorf("[%s] maxArea() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
