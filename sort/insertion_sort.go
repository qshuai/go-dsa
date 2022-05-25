package sort

// InsertionSort 插入排序（原地排序算法；稳定排序；时间复杂度O(n^2)）
// 基本思路：遍历每一个元素，如果该元素比上一个元素小，相邻元素依次交换，
// 直到该元素不再比上一个元素大
func InsertionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] <= arr[j+1] {
				break
			}

			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}
}
