package search

import (
	"testing"
)

func TestIsHappyNum(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "a valid happy number",
			num:  19,
			want: true,
		},
		{
			name: "invalid happy number(infinite loop if considering duplicate)",
			num:  11,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHappyNum(tt.num); tt.want != got {
				t.Errorf("IsHappyNum() = %t, want: %t", got, tt.want)
			}
		})
	}
}

func TestGetSquareSum(t *testing.T) {
	tests := []struct {
		name string
		num  int
		ret  int
	}{
		{
			name: "number without 0",
			num:  12345,
			ret:  55,
		},
		{
			name: "number suffixed with 0",
			num:  21000,
			ret:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSquareSum(tt.num); got != tt.ret {
				t.Errorf("getSquareSum() = %d, want: %d ", got, tt.ret)
			}
		})
	}
}
