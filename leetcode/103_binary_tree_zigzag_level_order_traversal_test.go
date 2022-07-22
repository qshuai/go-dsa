package leetcode

import (
	"reflect"
	"testing"

	"github.com/qshuai/go-dsa/types"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		args *types.Node
		want [][]int
	}{
		{
			name: "case-1",
			args: &types.Node{
				Value: 3,
				Left: &types.Node{
					Value: 9,
				},
				Right: &types.Node{
					Value: 20,
					Left: &types.Node{
						Value: 15,
					},
					Right: &types.Node{
						Value: 7,
					},
				},
			},
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
