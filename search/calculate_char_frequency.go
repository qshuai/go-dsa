package search

// CalcCharCntWithCaseSensitive 一个字符串由26个英文字母组成，计算字符串中各个字符出现的次数[注意区分大小写]
func CalcCharCntWithCaseSensitive(str string) []int {
	if len(str) <= 0 {
		return nil
	}

	res := make([]int, 52)
	for _, c := range str {
		if c < 'a' {
			// Uppercase
			// +26将大写字母放在后面
			res[c-'A'+26]++
		} else {
			res[c-'a']++
		}
	}

	return res
}

// CalcCharCntWithCaseInsensitive 一个字符串由26个英文字母组成，计算字符串中各个字符出现的次数[注意不区分大小写]
func CalcCharCntWithCaseInsensitive(str string) []int {
	if len(str) <= 0 {
		return nil
	}

	res := make([]int, 26)
	for _, c := range str {
		if c < 'a' {
			res[c-'A']++
		} else {
			res[c-'a']++
		}
	}

	return res
}
