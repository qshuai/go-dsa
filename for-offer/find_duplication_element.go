package for_offer

// Question:
//   在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。数组中某些数字是重复的，
//   但不知道有几个数字是重复的，也不知道每个数字重复几次。请找出数组中任意一个重复的数字。

// O(n) + O(1): 时间复杂度为O(n), 空间复杂度为O(1)
// 假设一种极端情况，在第一次外层循环过程中，内层循环每次都会触发[array[i] != i], 导致
// 内层循环已经将数组排完序了; 这样外层循环接下来的都不会触发内部循环，这样的循环次数为2n，
// 时间复杂度为O(n).
func findDuplicationElement(array []int) int {
	for i := 0; i < len(array); i++ {
		for array[i] != i {
			// the condition of stopping travelling
			if array[array[i]] == array[i] {
				return array[i]
			}

			// swap
			array[array[i]], array[i] = array[i], array[array[i]]
		}
	}

	// -1 represents duplicated element not found
	return -1
}
