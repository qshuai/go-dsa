package leetcode

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty array",
			args: args{
				nums: nil,
			},
			want: nil,
		},
		{
			name: "single element array",
			args: args{
				nums: []int{3},
			},
			want: []int{3},
		},
		{
			name: "case1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "case2",
			args: args{
				nums: []int{1, 1, 1, 3, 12},
			},
			want: []int{1, 1, 1, 3, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)

			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() = %v but want: %v", tt.args.nums, tt.want)
			}
		})
	}
}
