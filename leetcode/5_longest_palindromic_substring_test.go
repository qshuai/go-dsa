package leetcode

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "case-1",
			args: "babad",
			want: "bab",
		},
		{
			name: "case-2",
			args: "cbbd",
			want: "bb",
		},
		{
			name: "case-3",
			args: "a",
			want: "a",
		},
		{
			name: "case-4",
			args: "ac",
			want: "a",
		},
		{
			name: "case-5",
			args: "ccc",
			want: "ccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
