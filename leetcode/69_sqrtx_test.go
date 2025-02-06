package leetcode

import "testing"

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "case-1",
			x:    4,
			want: 2,
		},
		{
			name: "case-2",
			x:    8,
			want: 2,
		},
		{
			name: "case-3",
			x:    7,
			want: 2,
		},
		{
			name: "case-4",
			x:    1,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("[%s] mySqrt() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
