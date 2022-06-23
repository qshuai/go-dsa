package leetcode

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/restore-ip-addresses/

// restoreIpAddresses Given a string s containing only digits, return all possible valid IP addresses
// that can be formed by inserting dots into s. You are not allowed to reorder or remove any digits
// in s. You may return the valid IP addresses in any order.
//
// Constraints:
// 1 <= s.length <= 20
// s consists of digits only.
func restoreIpAddresses(s string) []string {
	collector, state := make([]string, 0), make([]string, 0)
	travelIP(s, &state, &collector)

	return collector
}

func travelIP(tail string, state, collector *[]string) {
	const MaxDigitsOfIpSegment = 3
	const IpSegmentCnt = 4
	for i := 0; i < MaxDigitsOfIpSegment && i < len(tail); i++ {
		if ok := validIpSegment(tail[:i+1]); !ok {
			break
		}

		// 1. 保存状态
		*state = append(*state, tail[:i+1])
		if i == len(tail)-1 && len(*state) == IpSegmentCnt {
			// 遍历完字符串，并且ip段数为4个，就代表找到了一个合法的ip
			*collector = append(*collector, strings.Join(*state, "."))
		}

		// 2. 深度优先搜索
		travelIP(tail[i+1:], state, collector)

		// 3. 恢复状态
		*state = (*state)[:len(*state)-1]
		if len(*state) >= IpSegmentCnt {
			break
		}
	}
}

func validIpSegment(seg string) bool {
	if len(seg) <= 0 || len(seg) > 3 {
		return false
	}
	if len(seg) == 1 {
		return true
	}

	if seg[0] == '0' {
		// 前导零是非法的字段
		return false
	}

	const MaxSegVal = 255
	segVal, err := strconv.Atoi(seg)
	if err != nil {
		panic(err)
	}

	return segVal <= MaxSegVal
}
