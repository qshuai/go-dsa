package leetcode

import "testing"

func Test_trap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "case-1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "case-2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("[%s] trap() = %v, want %v", tt.name, got, tt.want)
			}

			if got := trap2(tt.height); got != tt.want {
				t.Errorf("[%s] trap2() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
