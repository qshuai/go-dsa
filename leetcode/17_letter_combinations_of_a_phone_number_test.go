package leetcode

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "two letters",
			args: "23",
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
