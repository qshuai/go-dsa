package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func Test_searchRankK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case-1",
			args: args{
				arr: []int{2, 3, 1, 6, 2, 4},
				k:   3,
			},
			want: 3,
		},
		{
			name: "case-2",
			args: args{
				arr: []int{2, 3, 1, 6, 2, 4},
				k:   4,
			},
			want: 2,
		},
		{
			name: "case-3",
			args: args{
				arr: []int{2, 3, 1, 6, 2, 4},
				k:   5,
			},
			want: 2,
		},
		{
			name: "case-4",
			args: args{
				arr: []int{-1, -1},
				k:   2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(utils.Copy(tt.args.arr), tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}

			if got := findKthLargest2(utils.Copy(tt.args.arr), tt.args.k); got != tt.want {
				t.Errorf("findKthLargest2() = %v, want %v", got, tt.want)
			}
		})
	}
}
