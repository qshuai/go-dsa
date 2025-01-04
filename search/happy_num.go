package search

import (
	"github.com/qshuai/go-dsa/types"
)

// IsHappyNum 一个数字按照十进制的位，求各个位的平方和，如果平方和为1这位happy num，
// 不为1则对平方和再次进行各位求平方和, 依次进行. 为了防止出现循环的情况
// 需要一个set来保存之前平方和不为1的数字。
func IsHappyNum(num int) bool {
	set := types.NewSet[int]()
	for num != 1 {
		if set.Has(num) {
			return false
		}

		set.Add(num)

		num = getSquareSum(num)
	}

	return true
}

func getSquareSum(num int) int {
	var sum int
	for num > 0 {
		// get low bit square sum
		sum += (num % 10) * (num % 10)
		num /= 10
	}

	return sum
}
