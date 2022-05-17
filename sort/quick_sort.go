package sort

// QuickSort 快速排序，原地排序算法，不稳定的排序算法
// 时间复杂度为nlogn（最坏的时间复杂度为n^2，取决于pivot的选取是否尽可能将数组平分）
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := partition(arr)
	QuickSort(arr[:pivot])
	QuickSort(arr[pivot+1:])
}

func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	var i int
	for j := 0; j < len(arr) - 1; j++ {
		if arr[j] >= pivot {
			continue
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
	}

	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}
