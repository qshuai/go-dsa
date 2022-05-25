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
		Name:     "Truffle array",
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

// GetTestCases 获取排序所需的测试case。采用深拷贝方式，以防修改测试数据，影响后续测试
func GetTestCases() []TestItem {
	ret := make([]TestItem, 0, len(cases))
	for _, item := range cases {
		ret = append(ret, TestItem{
			Name:     item.Name,
			Args:     copyArray(item.Args),
			Min:      item.Min,
			Max:      item.Max,
			Expected: copyArray(item.Expected),
		})
	}

	return ret
}

func copyArray[T any](src []T) []T {
	if src == nil {
		return nil
	}
	if len(src) <= 0 {
		return []T{}
	}

	ret := make([]T, len(src))
	copy(ret, src)
	return ret
}
