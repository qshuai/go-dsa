package leetcode

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "valid parentheses",
			args: "(){}",
			want: true,
		},
		{
			name: "nested parentheses",
			args: "{()}",
			want: true,
		},
		{
			name: "invalid parentheses",
			args: "(]",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
