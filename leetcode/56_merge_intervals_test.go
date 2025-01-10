package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case-1",
			args: args{
				intervals: [][]int{{8, 10}, {2, 6}, {15, 18}, {1, 3}},
			},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "case-2",
			args: args{
				intervals: [][]int{{1, 4}, {4, 5}},
			},
			want: [][]int{{1, 5}},
		},
		{
			name: "case-3",
			args: args{
				intervals: [][]int{{1, 5}, {1, 4}},
			},
			want: [][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
