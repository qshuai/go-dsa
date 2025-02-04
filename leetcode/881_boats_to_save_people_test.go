package leetcode

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case-1",
			args: args{
				people: []int{1, 2},
				limit:  3,
			},
			want: 1,
		},
		{
			name: "case-2",
			args: args{
				people: []int{3, 2, 2, 1},
				limit:  3,
			},
			want: 3,
		},
		{
			name: "case-3",
			args: args{
				people: []int{3, 5, 3, 4},
				limit:  4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.args.people, tt.args.limit); got != tt.want {
				t.Errorf("[%s] numRescueBoats() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
