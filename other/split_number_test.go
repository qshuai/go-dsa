package other

import (
	"reflect"
	"testing"
)

func TestGetNumBits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		bits int
	}{
		{
			name: "zero",
			num:  0,
			bits: 1,
		},
		{
			name: "negative number",
			num:  -233,
			bits: 3,
		},
		{
			name: "normal positive number",
			num:  23399,
			bits: 5,
		},
		{
			name: "number suffixed with several zero",
			num:  20000,
			bits: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumBits(tt.num); got != tt.bits {
				t.Errorf("GetNumBits() = %d, want: %d", got, tt.bits)
			}
		})
	}
}

func TestSplitNum(t *testing.T) {
	tests := []struct {
		name string
		num  int
		res  []int
	}{
		{
			name: "normal positive number",
			num:  123456,
			res:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "normal negative number",
			num:  -73827,
			res:  []int{7, 3, 8, 2, 7},
		},
		{
			name: "number suffixed with zero",
			num:  20000,
			res:  []int{2, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitNum(tt.num); !reflect.DeepEqual(got, tt.res) {
				t.Errorf("SplitNum() = %v, want: %v", got, tt.res)
			}
		})
	}
}
