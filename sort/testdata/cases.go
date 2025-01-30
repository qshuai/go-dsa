package testdata

import (
	"math"
)

type TestItem struct {
	Name string
	Args []int
	Min  int
	Max  int

	Expected []int
}

var cases = []TestItem{
	{
		Name:     "Nil array",
		Args:     nil,
		Min:      math.MinInt,
		Max:      math.MaxInt,
		Expected: nil,
	},
	{
		Name:     "Empty array",
		Args:     []int{},
		Min:      math.MinInt,
		Max:      math.MaxInt,
		Expected: []int{},
	},
	{
		Name:     "Shuffled array",
		Args:     []int{2, 4, 8, 3, 1, 6, 5, 7},
		Min:      1,
		Max:      8,
		Expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		Name:     "Array with negative integer",
		Args:     []int{-1, 1, 2, 1, 3},
		Min:      -1,
		Max:      3,
		Expected: []int{-1, 1, 1, 2, 3},
	},
	{
		Name:     "Sorted array",
		Args:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Min:      1,
		Max:      10,
		Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		Name:     "Reversed array",
		Args:     []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		Min:      1,
		Max:      10,
		Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		Name:     "One-step sorting",
		Args:     []int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10},
		Min:      1,
		Max:      10,
		Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

// GetTestCases 获取排序所需的测试cases
func GetTestCases() []TestItem {
	return cases
}
