package tip

// tips:
// num % 10 will get the last bit number
// num / 10 will get a new number the last bit lost

func SplitNum(num int) []int {
	// convert a positive number
	if num < 0 {
		num = 0 - num
	}

	bits := getNumBits(num)
	ret := make([]int, bits)

	for num > 0 {
		bits--
		ret[bits] = num % 10
		num /= 10
	}

	return ret
}

func getNumBits(num int) int {
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
