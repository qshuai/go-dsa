package search

func binarySearchNoRecursion(array []int, target int) int {
	start := 0
	end := len(array) - 1
	var mid int

	for start <= end {
		mid = (start + end) / 2
		if array[mid] > target {
			end = mid - 1
		} else if array[mid] < target {
			start = mid + 1
		} else {
			return mid
		}

	}

	return -1
}

func binarySearchWithRecursion(array []int, start, end, target int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2
	if array[mid] < target {
		return binarySearchWithRecursion(array, mid+1, end, target)
	} else if array[mid] > target {
		return binarySearchWithRecursion(array, start, mid-1, target)
	} else {
		return mid
	}
}
