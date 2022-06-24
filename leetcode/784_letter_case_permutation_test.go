package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func Test_letterCasePermutation(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "case-1",
			args: "a1b2",
			want: []string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			name: "case-2",
			args: "3z4",
			want: []string{"3z4", "3Z4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCasePermutation(tt.args); !utils.ElementEqual(got, tt.want) {
				t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
