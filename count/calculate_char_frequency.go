package count

// 一个字符串由26个小写字母组成，计算字符串中各个字符出现的次数
func calculateCharFrequency(str string) []int {
	ret := make([]int, 26)

	for _, c := range str {
		ret[c-'a']++
	}

	return ret
}

// 一个字符串由26个英文字母组成，计算字符串中各个字符出现的次数[注意区分大小写]
func calculateCharFrequency2(str string) []int {
	ret := make([]int, 52)

	for _, c := range str {
		// upper case char
		if c < 'a' {
			ret[c-'A'+26]++
		} else {
			ret[c-'a']++
		}
	}

	return ret
}

// 一个字符串由26个英文字母组成，计算字符串中各个字符出现的次数[注意不区分大小写]
func calculateCharFrequency3(str string) []int {
	ret := make([]int, 26)

	for _, c := range str {
		// upper case char
		if c < 'a' {
			ret[c-'A']++
		} else {
			ret[c-'a']++
		}
	}

	return ret
}

// 可选方案，但是没有上面的方案效率高。 以26个小写英文字母组成的字符串为例:
func alternativeOptionWithHashMap(str string) map[rune]int {
	ret := make(map[rune]int)

	for _, c := range str {
		ret[c]++
	}

	return ret
}
