package search

import "testing"

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
}

func TestIndexOfUsingBF(t *testing.T) {
	for _, tt := range indexOfCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOfUsingBF(tt.args.str, tt.args.sub); got != tt.want {
				t.Errorf("IndexOfUsingBF() = %v, want %v", got, tt.want)
			}

			if got := IndexOfUsingBM(tt.args.str, tt.args.sub); got != tt.want {
				t.Errorf("IndexOfUsingBM() = %v, want %v", got, tt.want)
			}
		})
	}
}
