package sort

// QuickSort 快速排序
func QuickSort(arr []int) {
	if len(arr) <= 0 {
		return
	}

	pivot := partition(arr)
	QuickSort(arr[:pivot])
	QuickSort(arr[pivot+1:])
}

func partition(arr []int) int {
	low := 0
	high := len(arr) - 1

	for {
		if low >= high {
			break
		}

		for low < high && arr[low] < arr[high] {
			low++
		}
		arr[low], arr[high] = arr[high], arr[low]

		for low < high && arr[low] < arr[high] {
			high--
		}
		arr[low], arr[high] = arr[high], arr[low]
	}

	return low
}
