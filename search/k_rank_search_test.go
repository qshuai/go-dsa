package search

import (
	"testing"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRankK(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("searchRankK() = %v, want %v", got, tt.want)
			}
		})
	}
}
