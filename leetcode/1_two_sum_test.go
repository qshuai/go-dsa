package leetcode

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		target int
		array  []int
		res    []int
	}{
		{
			name:   "not found",
			target: 10,
			array:  []int{1, 2, 3, 4, 5},
			res:    nil,
		},
		{
			name:   "exits",
			target: 8,
			array:  []int{1, 2, 3, 4, 5},
			res:    []int{2, 4},
		},
		{
			name:   "exits",
			target: 6,
			array:  []int{3, 2, 4},
			res:    []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := twoSum(tt.array, tt.target); !utils.ElementEqual(tt.res, result) {
				t.Errorf("twoSum() = %v want %v", result, tt.res)
			}

			if result := twoSumN2(tt.array, tt.target); !utils.ElementEqual(tt.res, result) {
				t.Errorf("twoSumN2() = %v want %v", result, tt.res)
			}
		})
	}
}
