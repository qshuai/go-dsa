package leetcode

import "testing"

func Test_judgePoint24(t *testing.T) {
	type args struct {
		cards []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case-1",
			args: args{
				cards: []int{4, 1, 8, 7},
			},
			want: true,
		},
		{
			name: "case-2",
			args: args{
				cards: []int{1, 2, 1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgePoint24(tt.args.cards); got != tt.want {
				t.Errorf("judgePoint24() = %v, want %v", got, tt.want)
			}
		})
	}
}
