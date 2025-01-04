package utils

import (
	"testing"
)

func Test_bitCount(t *testing.T) {
	tests := []struct {
		name string
		args uint32
		want int
	}{
		{
			name: "case-1",
			args: 4,
			want: 1,
		},
		{
			name: "case-2",
			args: 3,
			want: 2,
		},
		{
			name: "case-3",
			args: 11,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BitCountUint32(tt.args); got != tt.want {
				t.Errorf("bitCount() = %v, want %v", got, tt.want)
			}

			if got := BitCountUint64(uint64(tt.args)); got != tt.want {
				t.Errorf("bitCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
