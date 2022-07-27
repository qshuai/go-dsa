package leetcode

import (
	"testing"
)

func Test_validSquare(t *testing.T) {
	type args struct {
		p1 []int
		p2 []int
		p3 []int
		p4 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case-1",
			args: args{
				p1: []int{0, 0},
				p2: []int{1, 1},
				p3: []int{1, 0},
				p4: []int{0, 1},
			},
			want: true,
		},
		{
			name: "case-2",
			args: args{
				p1: []int{0, 0},
				p2: []int{1, 1},
				p3: []int{1, 0},
				p4: []int{0, 12},
			},
			want: false,
		},
		{
			name: "case-3",
			args: args{
				p1: []int{1, 0},
				p2: []int{-1, 0},
				p3: []int{0, 1},
				p4: []int{0, -1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validSquare(tt.args.p1, tt.args.p2, tt.args.p3, tt.args.p4); got != tt.want {
				t.Errorf("validSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
