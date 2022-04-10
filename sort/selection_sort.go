package sort

// selectSort 选择排序（不稳定排序；时间复杂度O(n^2)）
// 这个算法的优点是避免了不必要的元素交换
func selectionSort(arr []int) {
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
