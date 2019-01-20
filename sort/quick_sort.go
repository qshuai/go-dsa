package sort

func QuickSort(array []int) []int {
	low := 0
	high := len(array) - 1

	if low < high {
		pivot := Partition(array)

		QuickSort(array[:pivot])
		QuickSort(array[pivot+1:])
	}

	return array
}

func Partition(array []int) int {
	low := 0
	high := len(array) - 1
	pivotKey := array[low]

	for low < high {
		for low < high && pivotKey < array[high] {
			high--
		}
		array[low], array[high] = array[high], array[low]

		for low < high && array[low] <= pivotKey {
			low++
		}
		array[low], array[high] = array[high], array[low]
	}

	return low
}
