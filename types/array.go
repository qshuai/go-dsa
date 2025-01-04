package types

// moveArray 数组右移K次，生成循环数组
// 原数组<code> [1, 2, 3, 4, 5, 6, 7]</code> 右移3次后结果为 [5,6,7,1,2,3,4]
// 基本思路：不开辟新的数组空间的情况下考虑在原数组上进行操作
// 1 将数组倒置，这样后k个元素就跑到了数组的前面，然后反转一下即可
// 2 同理后 len-k个元素只需要翻转就完成数组的k次移动
func moveArray(arr []int, k int) {
	if len(arr) == 0 || k%len(arr) == 0 {
		return
	}

	// k有可能超过数组长度，而数组长度整数倍的移动仍为数组本身，故只需要考虑多出的移动步数
	k = k % len(arr)

	// reverse source array
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	// reverse the first k elements
	for i := 0; i < k/2; i++ {
		arr[i], arr[k-1-i] = arr[k-1-i], arr[i]
	}

	// reverse the last len(arr) - k elements
	for i := k; i < (k+len(arr))/2; i++ {
		arr[i], arr[len(arr)-1-i+k] = arr[len(arr)-1-i+k], arr[i]
	}
}
