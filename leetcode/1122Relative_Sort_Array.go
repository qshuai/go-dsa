package leetcode

// https://leetcode.com/problems/relative-sort-array/

// relativeSortArray
//
// Constraints:
// 1 <= arr1.length, arr2.length <= 1000
// 0 <= arr1[i], arr2[i] <= 1000
// All the elements of arr2 are distinct.
// Each arr2[i] is in arr1.
func relativeSortArray(arr1 []int, arr2 []int) []int {
	min, max := arr1[0], arr1[0]
	for _, num := range arr1 {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	buckets := make([]int, max-min+1)
	for _, num := range arr1 {
		buckets[num-min]++
	}

	var offset int
	for _, num := range arr2 {
		for buckets[num-min] > 0 {
			arr1[offset] = num
			buckets[num-min]--
			offset++
		}
	}

	for idx, item := range buckets {
		for item > 0 {
			arr1[offset] = min + idx
			item--
			offset++
		}
	}

	return arr1
}
