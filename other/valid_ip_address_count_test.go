package other

import (
	"reflect"
	"testing"
)

func Test_validIpAddressCount(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "case-1",
			args: "25525511135",
			want: 2,
		},
		{
			name: "case-2",
			args: "0000",
			want: 1,
		},
		// +---+---+---+---+---+---+---+---+
		// |   |   | 1 | 0 | 1 | 0 | 2 | 3 |
		// +---+---+---+---+---+---+---+---+
		// | 0 | 1 | 0 | 0 | 0 | 0 | 0 | 0 |
		// | 1 | 0 | 1 | 1 | 1 | 0 | 0 | 0 |
		// | 2 | 0 | 0 | 1 | 1 | 2 | 1 | 0 |
		// | 3 | 0 | 0 | 0 | 1 | 2 | 3 | 3 |
		// | 4 | 0 | 0 | 0 | 0 | 1 | 2 | 5 |
		// +---+---+---+---+---+---+---+---+
		{
			name: "case-3",
			args: "101023",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIpAddressCount(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validIpAddressCount() = %d, want %d", got, tt.want)
			}
		})
	}
}
