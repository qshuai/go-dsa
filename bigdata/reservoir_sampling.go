package bigdata

import (
	"math/rand"
	"time"
)

// sampling 蓄水池算法的go实现，目的：从一个超大数组中随机采样k个样本
// 时间复杂度: O(n)
func sampling(data *[]int, k int) []int {
	if data == nil {
		return nil
	}
	if len(*data) <= k {
		return *data
	}

	res := make([]int, 0, k)
	elements := *data
	for i := 0; i < k; i++ {
		res = append(res, elements[i])
	}

	source := rand.NewSource(time.Now().UnixMilli())
	r := rand.New(source)
	for i := k; i < len(elements); i++ {
		idx := r.Int63n(int64(i) + 1)
		if idx < int64(k) {
			res[idx] = elements[i]
		}
	}

	return res
}
