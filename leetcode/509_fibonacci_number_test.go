package leetcode

import "testing"

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "case-1",
			args: 0,
			want: 0,
		},
		{
			name: "case-2",
			args: 1,
			want: 1,
		},
		{
			name: "case-3",
			args: 3,
			want: 2,
		},
		{
			name: "case-4",
			args: 4,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
			if got := fibLoop(tt.args); got != tt.want {
				t.Errorf("fibLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
