package bigdata

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestSampling(t *testing.T) {
	// 准备一个大数组
	n := 100000
	k := 1000
	samples := sampling(arange(0, n, 1), k)

	// 检查样本的分布情况
	bucket := 5000
	hist(samples, 0, bucket)
}

func arange(start, end, step int) []int {
	if start > end {
		return nil
	}
	if step <= 0 {
		return nil
	}
	if start == end {
		return []int{start}
	}

	length := (end - start) / step
	if (end-start)%step != 0 {
		length++
	}
	res := make([]int, 0, length)
	for i := start; i < end; i += step {
		res = append(res, i)
	}

	return res
}

func hist(data []int, start, step int) {
	if len(data) <= 0 {
		return
	}

	m := make(map[int]int)
	keys := make([]int, 0)
	for _, v := range data {
		multi := (v - start) / step
		if _, ok := m[multi]; !ok {
			keys = append(keys, multi)
		}

		m[multi]++
	}

	sort.Ints(keys)
	for _, k := range keys {
		v := m[k]
		fmt.Printf("%3d: %s\n", k, strings.Repeat("+", v))
	}
}
