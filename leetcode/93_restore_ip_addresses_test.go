package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func Test_restoreIpAddresses(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "case-1",
			args: "25525511135",
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "case-2",
			args: "0000",
			want: []string{"0.0.0.0"},
		},
		{
			name: "case-3",
			args: "101023", // 递归树：https://github.com/qshuai/go-dsa/blob/master/images/restore_ip_addresses.png
			want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args); !utils.ElementEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
