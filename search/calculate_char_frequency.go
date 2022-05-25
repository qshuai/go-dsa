package search

// CalcCharCntWithCaseSensitive 一个字符串由26个英文字母组成，计算字符串中各个字符出现的次数[注意区分大小写]
func CalcCharCntWithCaseSensitive(str string) []int {
	if len(str) <= 0 {
		return nil
	}

	ret := make([]int, 52)
	for _, c := range str {
		if c < 'a' {
			// Uppercase
			// +26将大写字母放在后面
			ret[c-'A'+26]++
		} else {
			ret[c-'a']++
		}
	}

	return ret
}

// CalcCharCntWithCaseInsensitive 一个字符串由26个英文字母组成，计算字符串中各个字符出现的次数[注意不区分大小写]
func CalcCharCntWithCaseInsensitive(str string) []int {
	if len(str) <= 0 {
		return nil
	}

	ret := make([]int, 26)
	for _, c := range str {
		if c < 'a' {
			ret[c-'A']++
		} else {
			ret[c-'a']++
		}
	}

	return ret
}
