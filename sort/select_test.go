package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSelectSort(t *testing.T) {
	tests := []struct {
		name     string
		target   []int
		expected []int
	}{
		{
			name:     "case1(has sorted array)",
			target:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "case2(reversed array)",
			target:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "case3(truffle array)",
			target:   []int{2, 3, 1, 345, 2, 56, 86, 34, 0, -23},
			expected: []int{-23, 0, 1, 2, 2, 3, 34, 56, 86, 345},
		},
		{
			name:     "case4(one step sort)",
			target:   []int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range tests {
		sorted := selectSort(test.target)
		if !equal(sorted, test.expected) {
			t.Errorf("test: %s failed, expected: %v, but got: %v",
				test.name, test.expected, sorted)
		}
	}
}

func TestRandomArray(t *testing.T) {
	// random array sort testing
	arrayLength := 255
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 10000; i++ {
		array := make([]int, arrayLength)
		compare := make([]int, arrayLength)

		for j := 0; j < arrayLength; j++ {
			array[j] = seed.Int()
		}
		copy(compare, array)

		// compare group
		sort.Ints(compare)

		// implementation group
		calcRet := selectSort(array)

		if !equal(calcRet, compare) {
			t.Errorf("test: random array sort failed, expected: %v,"+
				" but got: %v", compare, calcRet)
		}
	}
}
