package utils

import "testing"

func TestElementEqual(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case-1",
			args: args{
				a: []string{"3z4", "3Z4"},
				b: []string{"3Z4", "3z4"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElementEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ElementEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
