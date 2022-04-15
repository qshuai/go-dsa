package search

// searchRankK 从大到小，找到排名第K的元素(值相同的元素，排名加1，也就是说排名是连续的[可以查看测试用例])
func searchRankK(arr []int, k int) int {
	if k > len(arr) || k <= 0 || len(arr) <= 0 {
		return -1
	}

	var offset int
	for {
		pivot := partition(arr)
		if pivot+1+offset == k {
			return arr[pivot]
		} else if pivot+1+offset > k {
			arr = arr[:pivot]
		} else {
			offset += pivot
			arr = arr[pivot:]
		}
	}
}

// 分区，大的元素放在前面（索引小的位置）
func partition(arr []int) int {
	i, j := 0, len(arr)-1
	for {
		if i >= j {
			break
		}

		for i < j && arr[i] >= arr[j] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
		for i < j && arr[i] >= arr[j] {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	return i
}
