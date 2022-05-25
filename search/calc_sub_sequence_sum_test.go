package search

import "testing"

func TestCalcMaxSumForSubSequence(t *testing.T) {
	tests := []struct {
		name     string
		args     []int
		expected int
	}{
		{
			name:     "all negative integer",
			args:     []int{-12, -31, -9, -45},
			expected: -9,
		},
		{
			name:     "normal data set",
			args:     []int{11, -20, 1, 2, 89, -45, -2, 4},
			expected: 92,
		},
		{
			name:     "last element is the result",
			args:     []int{1, 2, 3, -40, 40},
			expected: 40,
		},
		{
			name:     "across negative integer",
			args:     []int{10, -2, 12, -30, 2},
			expected: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := CalcMaxSumForSubSequence(tt.args)
			if ret != tt.expected {
				t.Errorf("%s, want: %d, but got: %d", tt.name, tt.expected, ret)
			}
		})
	}
}
