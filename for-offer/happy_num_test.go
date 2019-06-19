package for_offer

import (
	"testing"
)

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

	for _, test := range tests {
		result := getSquareSum(test.num)
		if result != test.ret {
			t.Errorf("test case: %s(target num: %d), want: %d but got: %d",
				test.name, test.num, test.ret, result)
		}
	}
}

func TestIsHappyNum(t *testing.T) {
	tests := []struct {
		name string
		num  int
		ret  bool
	}{
		{
			name: "a valid happy number",
			num:  19,
			ret:  true,
		},
		{
			name: "invalid happy number(infinite loop if considering duplicate)",
			num:  11,
			ret:  false,
		},
	}

	for _, test := range tests {
		result := isHappyNum(test.num)
		if test.ret != result {
			t.Errorf("test case: %s(target num: %d), want: %t but got: %t",
				test.name, test.num, test.ret, result)
		}
	}
}
