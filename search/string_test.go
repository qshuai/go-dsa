package search

import (
	"reflect"
	"testing"
)

type args struct {
	str string
	sub string
}

var indexOfCases = []struct {
	name string
	args args
	want int
}{
	{
		name: "case-1",
		args: args{
			str: "abc",
			sub: "bc",
		},
		want: 1,
	},
	{
		name: "case-2",
		args: args{
			str: "abcd",
			sub: "bc",
		},
		want: 1,
	},
	{
		name: "case-3",
		args: args{
			str: "abcd",
			sub: "bd",
		},
		want: -1,
	},
	{
		name: "case-4",
		args: args{
			str: "abcd",
			sub: "abcd",
		},
		want: 0,
	},
	{
		name: "case-5",
		args: args{
			str: "abcde",
			sub: "abcdex",
		},
		want: -1,
	},
	{
		name: "suffix",
		args: args{
			str: "yyyybdxbdyyy",
			sub: "ybdxbd",
		},
		want: 3,
	},
	{
		name: "prefix",
		args: args{
			str: "yyyybdxbdyyy",
			sub: "bdxbd",
		},
		want: 4,
	},
	{
		name: "end matched",
		args: args{
			str: "xxxabc",
			sub: "abc",
		},
		want: 3,
	},
	{
		name: "KMP",
		args: args{
			str: "abababcaa",
			sub: "ababc",
		},
		want: 2,
	},
}

func TestIndexOfUsingBF(t *testing.T) {
	for _, tt := range indexOfCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOfUsingBF(tt.args.str, tt.args.sub); got != tt.want {
				t.Errorf("IndexOfUsingBF() = %d, want %d", got, tt.want)
			}

			if got := IndexOfUsingBM(tt.args.str, tt.args.sub); got != tt.want {
				t.Errorf("IndexOfUsingBM() = %d, want %d", got, tt.want)
			}

			if got := IndexOfUsingRK(tt.args.str, tt.args.sub); got != tt.want {
				t.Errorf("IndexOfUsingRK() = %d, want %d", got, tt.want)
			}

			if got := IndexOfUsingKMP(tt.args.str, tt.args.sub); got != tt.want {
				t.Errorf("IndexOfUsingKMP = %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_buildNextArray(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []int
	}{
		{
			name: "case-1",
			args: "abacabab",
			want: []int{0, 0, 1, 0, 1, 2, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildNextArray(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildNextArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
