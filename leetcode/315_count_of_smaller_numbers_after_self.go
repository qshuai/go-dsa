package leetcode

// https://leetcode.com/problems/count-of-smaller-numbers-after-self/

// You are given an integer array nums and you have to return a new counts array.
// The counts array has the property where counts[i] is the number of smaller
// elements to the right of nums[i].
//
// Constraints:
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
func countSmaller(nums []int) []int {
	if len(nums) <= 0 {
		return nil
	}

	ni := make([]*NumIndex, len(nums))
	for i, num := range nums {
		ni[i] = &NumIndex{
			num: num,
			idx: i,
		}

		nums[i] = 0
	}

	mergeSort(ni, nums, 0, len(nums)-1)
	return nums
}

func mergeSort(ni []*NumIndex, nums []int, begin, end int) {
	if begin >= end {
		return
	}

	mid := begin + (end-begin)>>1
	mergeSort(ni, nums, begin, mid)
	mergeSort(ni, nums, mid+1, end)
	merge(ni, nums, begin, mid, end)
}

func merge(ni []*NumIndex, nums []int, begin, mid, end int) {
	i, j := begin, mid+1
	tmp := make([]*NumIndex, 0, end-begin+1)
	for i <= mid && j <= end {
		if ni[i].num <= ni[j].num {
			tmp = append(tmp, ni[j])
			j++
		} else {
			nums[ni[i].idx] += end - j + 1
			tmp = append(tmp, ni[i])
			i++
		}
	}
	for i <= mid {
		tmp = append(tmp, ni[i])
		i++
	}
	for j <= end {
		tmp = append(tmp, ni[j])
		j++
	}

	for i := begin; i <= end; i++ {
		ni[i] = tmp[i-begin]
		j++
	}
}

type NumIndex struct {
	num int
	idx int
}
