package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "case-1",
			args: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			name: "case-2",
			args: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			name: "case-3",
			args: []int{7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			name: "case-4",
			args: []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}

			if got := lengthOfLIS2(tt.args); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
