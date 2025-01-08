package leetcode

import (
	"testing"
)

func Test_balancedString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case-1",
			args: args{
				s: "QWER",
			},
			want: 0,
		},
		{
			name: "case-2",
			args: args{
				s: "QQWE",
			},
			want: 1,
		},
		{
			name: "case-3",
			args: args{
				s: "QQQW",
			},
			want: 2,
		},
		{
			name: "case-4",
			args: args{
				s: "WQWRQQQW",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedString(tt.args.s); got != tt.want {
				t.Errorf("balancedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
