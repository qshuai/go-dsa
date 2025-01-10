package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func Test_sortArrayByParityII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case-1",
			args: args{
				nums: []int{4, 2, 5, 7},
			},
			want: [][]int{{4, 7, 2, 5}, {2, 5, 4, 7}, {2, 7, 4, 5}},
		},
		{
			name: "case-2",
			args: args{
				nums: []int{2, 3},
			},
			want: [][]int{{2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := utils.Copy(tt.args.nums)
			if got := sortArrayByParityII(arr); !utils.ContainEqual(tt.want, got) {
				t.Errorf("sortArrayByParityII() = %v, want one of: %v", got, tt.want)
			}

			arr = utils.Copy(tt.args.nums)
			if got := sortArrayByParityII2(arr); !utils.ContainEqual(tt.want, got) {
				t.Errorf("sortArrayByParityII() = %v, want one of: %v", got, tt.want)
			}
		})
	}
}
