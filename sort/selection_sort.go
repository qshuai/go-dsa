package sort

// SelectionSort 选择排序（不稳定排序；时间复杂度O(n^2)）；优点是避免了不必要的元素交换
// 基本思路：遍历每一个元素，在该元素以及右侧寻找最小元素，如果找到了和当前元素交换位置，
// 直到遍历完整个数组
func SelectionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr); i++ {
		min := arr[i]
		target := -1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				target = j
			}
		}

		if target > 0 {
			arr[i], arr[target] = arr[target], arr[i]
		}
	}
}
