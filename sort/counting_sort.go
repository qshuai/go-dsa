package sort

// CountingSort 计数排序
// The data table generated by the code at the tail of this file.
// +-------------------------------+---+---+---+---+---+---+
// |             INDEX             | 0 | 1 | 2 | 3 | 4 | 5 |
// +-------------------------------+---+---+---+---+---+---+
// |              unsorted numbers | 3 | 2 | 1 | 0 | 2 | 2 |
// |                counting array | 1 | 1 | 3 | 1 |   |   |
// |              cumulative count | 1 | 2 | 5 | 6 |   |   |
// | Iteration in reverse order[2] |   |   |   |   | 2 |   |
// |              cumulative count | 1 | 2 | 4 | 6 |   |   |
// | Iteration in reverse order[2] |   |   |   | 2 | 2 |   |
// |              cumulative count | 1 | 2 | 3 | 6 |   |   |
// | Iteration in reverse order[0] | 0 |   |   | 2 | 2 |   |
// |              cumulative count | 0 | 2 | 3 | 6 |   |   |
// | Iteration in reverse order[1] | 0 | 1 |   | 2 | 2 |   |
// |              cumulative count | 0 | 1 | 3 | 6 |   |   |
// | Iteration in reverse order[2] | 0 | 1 | 2 | 2 | 2 |   |
// |              cumulative count | 0 | 1 | 2 | 6 |   |   |
// | Iteration in reverse order[3] | 0 | 1 | 2 | 2 | 2 | 3 |
// |              cumulative count | 0 | 1 | 2 | 5 |   |   |
// +-------------------------------+---+---+---+---+---+---+
func CountingSort(nums []int, min, max int) {
	counting := make([]int, max-min+1)
	for _, num := range nums {
		counting[num-min]++
	}

	// cumulative count
	var sum int
	for idx := range counting {
		counting[idx] += sum
		sum = counting[idx]
	}

	// create a new array to store sorted integer
	res := make([]int, len(nums))
	var key int // 减少内存分配
	for i := len(nums) - 1; i >= 0; i-- {
		key = nums[i] - min // 避免多次计算
		// -1是将次数转化为索引（比如一个最小的数，他的出现的次数是1，
		// 即counting[key]为1，但是在数组中的索引为0）
		counting[key]--
		res[counting[key]] = nums[i]
	}

	// replace in place
	copy(nums, res)
}
