package leetcode

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "case-1",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'a',
			},
			want: 'c',
		},
		{
			name: "case-2",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'c',
			},
			want: 'f',
		},
		{
			name: "case-3",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'd',
			},
			want: 'f',
		},
		{
			name: "case-4",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'j',
			},
			want: 'c',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %c, want %c", got, tt.want)
			}
		})
	}
}
