package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []string
	}{
		{
			name: "one pair",
			args: 1,
			want: []string{"()"},
		},
		{
			name: "two pair",
			args: 2,
			want: []string{"(())", "()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
