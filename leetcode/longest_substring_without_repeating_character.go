package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return 1
		} else {
			return 2
		}
	}

	var max, prev int
	var flag bool
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			if !flag {
				if prev+1 > max {
					max = prev + 1
				}
			}

			flag = true
			prev = 0
		} else {
			if flag {
				prev = 2
				flag = false
			} else {
				prev += 1
			}
		}
	}

	return max
}
