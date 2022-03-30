package leetcode

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "empty string",
			args: "",
			want: true,
		},
		{
			name: "Palindrome with separator character",
			args: "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "negative case",
			args: "race a car",
			want: false,
		},
		{
			name: "non alphanumeric string",
			args: ", :__+",
			want: true,
		},
		{
			name: "single alphanumeric string",
			args: ",v; _",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args); got != tt.want {
				t.Errorf("method: isPalindrome, case: %s got %v, but want %v", tt.args, got, tt.want)
			}

			if got := isPalindrome2(tt.args); got != tt.want {
				t.Errorf("method: isPalindrome2, case: %s got %v, but want %v", tt.args, got, tt.want)
			}
		})
	}
}
