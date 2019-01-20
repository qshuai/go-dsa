package sort

func QuickSort(array []int) []int {
	QSort(array, 0, len(array)-1)

	return array
}

func QSort(array []int, low, high int) {
	if low < high {
		pivot := Partition(array, low, high)

		QSort(array, low, pivot)
		QSort(array, pivot+1, high)
	}
}

func Partition(array []int, low, high int) int {
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
