package sort

// mergeSort 合并排序算法
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	if len(left) <= 0 && len(right) <= 0 {
		return []int{}
	}

	// 创建临时slice只存放left部分，可以减少内存分配
	tmp := make([]int, len(left))
	copy(tmp, left)
	left = left[:0]

	var i, j int
	for ; i < len(tmp) && j < len(right); {
		// <=中的=条件，可以保证该排序是稳定排序（相等元素，左边的还在左边）
		if tmp[i] <= right[j] {
			left = append(left, tmp[i])
			i++
		} else {
			left = append(left, right[j])
			j++
		}
	}
	if i < len(tmp) {
		left = append(left, tmp[i:]...)
	}
	if j < len(right) {
		left = append(left, right[j:]...)
	}

	return left
}
