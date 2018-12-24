package sort

// 时间复杂度O(n^2)
func straightInsertSort(array []int) []int {
	var minimal int
	var j int

	// 每次比较两张相邻的牌，如果后面的小，就先把这个小牌的数值保存起来，
	// 然后把这张小牌之前的牌和这张小牌比较一次，如果比小牌大，就将
	// 大牌后移一位，直到遇到一个小于或等于这张小牌。
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			minimal = array[i]

			for j = i - 1; j >= 0 && array[j] > minimal; j-- {
				array[j+1] = array[j]
			}

			array[j+1] = minimal
		}
	}

	return array
}
