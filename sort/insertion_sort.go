package sort

// insertSort 插入排序（稳定排序；时间复杂度O(n^2)）
func insertionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] <= arr[j+1] {
				break
			}

			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}
}
