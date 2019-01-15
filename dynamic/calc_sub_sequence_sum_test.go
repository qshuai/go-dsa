package dynamic

import "testing"

func TestCalcMaxSumForSubSequence(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected int
	}{
		{
			name:     "all negative numbers",
			array:    []int{-12, -31, -9, -45},
			expected: -9,
		},
		{
			name:     "normal data set",
			array:    []int{11, -20, 1, 2, 89, -45, -2, 4},
			expected: 92,
		},
		{
			name:     "last element is the result",
			array:    []int{1, 2, 3, -40, 40},
			expected: 40,
		},
	}

	for _, test := range tests {
		ret := calcMaxSumForSubSequence(test.array)
		if ret != test.expected {
			t.Errorf("%s, want: %d, but got: %d",
				test.name, test.expected, ret)
		}
	}
}
