package tip

import (
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

	for _, test := range tests {
		bits := getNumBits(test.num)
		if bits != test.bits {
			t.Errorf("test case: %s(target num: %d), want: %d but got: %d",
				test.name, test.num, test.bits, bits)
		}
	}
}

func TestSplitNum(t *testing.T) {
	tests := []struct {
		name string
		num  int
		ret  []int
	}{
		{
			name: "normal positive number",
			num:  123456,
			ret:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "normal negative number",
			num:  -73827,
			ret:  []int{7, 3, 8, 2, 7},
		},
		{
			name: "number suffixed with zero",
			num:  20000,
			ret:  []int{2, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		result := SplitNum(test.num)
		if !equalSlice(result, test.ret) {
			t.Errorf("test case: %s(target num: %d), want: %v but got: %v",
				test.name, test.num, test.ret, result)
		}
	}
}

func equalSlice(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
