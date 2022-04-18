package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "case-1",
			args: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "case-2",
			args: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sortColors(tt.args); !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("sortColors() = %v, but want: %v", tt.args, tt.want)
			}
		})
	}
}
