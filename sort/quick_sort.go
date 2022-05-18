package sort

// QuickSort 快速排序，原地排序算法，不稳定的排序算法
// 时间复杂度为nlogn（最坏的时间复杂度为n^2，取决于pivot的选取是否尽可能将数组平分）
// 基本思路：选择一个pivot，通过移动元素，达到pivot左侧的都小于它，右侧的均大约等于它；
// 整个数组按照pivot切分后，对每个小数组再次进行上述操作，直到每个小数组只剩下1个元素，
// 整个数组将达成有序的状态
func QuickSort(arr []int, fn func([]int) int) {
	if len(arr) <= 1 {
		return
	}

	pivot := fn(arr)
	QuickSort(arr[:pivot], fn)
	QuickSort(arr[pivot+1:], fn)
}

// PartitionHead 选择第一个元素作为pivot，进行分区
func PartitionHead(arr []int) int {
	pivot := arr[0]
	i, j := len(arr)-1, len(arr)-1
	for ; j > 0; j-- {
		if arr[j] <= pivot {
			continue
		}

		arr[i], arr[j] = arr[j], arr[i]
		i--
	}

	arr[i], arr[0] = arr[0], arr[i]
	return i
}

// PartitionMiddle 选择中间位置的元素作为pivot，进行分区
func PartitionMiddle(arr []int) int {
	idx := len(arr) / 2
	pivot := arr[idx]
	var i int
	for j := 0; j < len(arr); j++ {
		if arr[j] >= pivot {
			continue
		}

		// track pivot element's index
		if arr[i] == pivot {
			idx = j
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
	}

	arr[i], arr[idx] = arr[idx], arr[i]
	return i
}

// PartitionTail 选择最后一个元素作为pivot，进行分区
func PartitionTail(arr []int) int {
	pivot := arr[len(arr)-1]
	var i int
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] >= pivot {
			continue
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
	}

	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}

// PartitionMedian 三数取中法
func PartitionMedian(arr []int) int {
	begin, middle, end := arr[0], arr[len(arr)/2], arr[len(arr)-1]
	if begin < middle {
		if end < begin {
			return PartitionHead(arr)
		} else if middle < end {
			return PartitionMiddle(arr)
		} else {
			return PartitionTail(arr)
		}
	} else {
		if begin < end {
			return PartitionHead(arr)
		} else if middle > end {
			return PartitionMiddle(arr)
		} else {
			return PartitionTail(arr)
		}
	}
}
