package other

// SplitNum 将一个十进制的数值按照高位到低位的顺序，分别获取各位上的数字
// num % 10 获取十进制的最后一位数字
// num / 10 去除十进制的最后一位数
func SplitNum(num int) []int {
	// convert a positive number
	if num < 0 {
		num = 0 - num
	}

	bits := GetNumBits(num)
	ret := make([]int, bits)

	for num > 0 {
		bits--
		ret[bits] = num % 10
		num /= 10
	}

	return ret
}

// GetNumBits 获取十进制数据有由几位组成的
func GetNumBits(num int) int {
	var bits int
	div := 10

	for {
		if num/div == 0 {
			return bits + 1
		} else {
			bits++
			div *= 10
		}
	}
}
