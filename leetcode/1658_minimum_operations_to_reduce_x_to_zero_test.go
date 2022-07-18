package leetcode

import (
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case-1",
			args: args{
				nums: []int{1, 1, 4, 2, 3},
				x:    5,
			},
			want: 2,
		},
		{
			name: "case-2",
			args: args{
				nums: []int{5, 6, 7, 8, 9},
				x:    4,
			},
			want: -1,
		},
		{
			name: "case-3",
			args: args{
				nums: []int{3, 2, 20, 1, 1, 3},
				x:    10,
			},
			want: 5,
		},
		{
			name: "case-4",
			args: args{
				nums: []int{8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904, 8819, 1231, 6309},
				x:    134365,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
