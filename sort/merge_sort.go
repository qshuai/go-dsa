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

func merge(a, b []int) []int {
	if len(a) <= 0 && len(b) <= 0 {
		return []int{}
	}

	ret := make([]int, 0, len(a)+len(b))
	var i, j int
	for ; i < len(a) && j < len(b); {
		if a[i] <= b[j] {
			ret = append(ret, a[i])
			i++
		} else {
			ret = append(ret, b[j])
			j++
		}
	}
	if i < len(a) {
		ret = append(ret, a[i:]...)
	}
	if j < len(b) {
		ret = append(ret, b[j:]...)
	}

	return ret
}
