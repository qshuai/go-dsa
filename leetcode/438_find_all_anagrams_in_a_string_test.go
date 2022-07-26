package leetcode

import (
	"reflect"
	"testing"
)

func Test_findanagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case-1",
			args: args{
				s: "cbaebabacd",
				p: "abc",
			},
			want: []int{0, 6},
		},
		{
			name: "case-2",
			args: args{
				s: "abab",
				p: "ab",
			},
			want: []int{0, 1, 2},
		},
		{
			name: "case-3",
			args: args{
				s: "baa",
				p: "aa",
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findanagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findanagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
