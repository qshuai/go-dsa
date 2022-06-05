package leetcode

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "case-1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "case-2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "case-3",
			s:    "pwwkew",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
