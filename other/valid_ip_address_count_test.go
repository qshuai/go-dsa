package other

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
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
		{
			name: "case-1",
			args: "101023",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %d, want %d", got, tt.want)
			}
		})
	}
}
