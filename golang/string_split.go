package golang

// SplitStringWithRepeatedChar 在由英文字母组成的字符串中，任何两个或两个以上的相同字符都是该字符串的分隔符，
// eg: "helloworld" => "he", "oworld"
func SplitStringWithRepeatedChar(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}
	if len(str) == 2 {
		if str[0] == str[1] {
			return nil
		} else {
			return []string{str}
		}
	}

	var res []string
	var flag bool
	var start, offset int

	for i := 1; i < len(str); i++ {
		if str[i-1] == str[i] {
			flag = true
			offset++
			continue
		}

		if flag {
			if start != i-offset-1 {
				res = append(res, str[start:i-offset-1])
			}
			flag = false
			start = i
			offset = 0
		}
	}

	if !flag {
		res = append(res, str[start:])
	} else {
		res = append(res, str[start:len(str)-offset-1])
	}

	return res
}

func SplitStringWithRepeatedChar2(str string) []string {
	var res []string

	var tmp string
	for i := 0; i < len(str); {
		j := i + 1
		for j < len(str) && str[j] == str[j-1] {
			j++
		}

		if j-i == 1 {
			tmp += string(str[i])
		} else {
			if tmp != "" {
				res = append(res, tmp)
			}
			tmp = ""
		}

		i = j
	}

	if tmp != "" {
		res = append(res, tmp)
	}

	return res
}
