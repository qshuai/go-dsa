package leetcode

// https://leetcode.com/problems/letter-case-permutation/

const (
	DigitBegin = '0'
	DigitEnd   = '9'

	LetterOfa = 'a'
	CaseDiff  = 'a' - 'A'
)

// Given a string s, you can transform every letter individually to be lowercase or
// uppercase to create another string. Return a list of all possible strings we could
// create. Return the output in any order.
//
// Constraints:
// 1 <= s.length <= 12
// s consists of lowercase English letters, uppercase English letters, and digits.
func letterCasePermutation(s string) []string {
	bs := []byte(s)
	collector := make([]string, 0)
	dfsLetter(bs, 0, &collector)
	return collector
}

// dfsLetter 深度优先遍历字符串
func dfsLetter(bs []byte, cursor int, collector *[]string) {
	if cursor >= len(bs) {
		// 已经遍历完成，直接收集结果
		*collector = append(*collector, string(bs))
		return
	}

	if bs[cursor] >= DigitBegin && bs[cursor] <= DigitEnd {
		// 遇到数字，略过且进入下次循环
		dfsLetter(bs, cursor+1, collector)
		return
	}

	// 不转换直接进行下次循环
	dfsLetter(bs, cursor+1, collector)

	// 转化该字符大小写并进入下次循环
	backup := bs[cursor]
	if bs[cursor] >= LetterOfa {
		bs[cursor] -= CaseDiff
	} else {
		bs[cursor] += CaseDiff
	}
	dfsLetter(bs, cursor+1, collector)
	// 恢复状态
	bs[cursor] = backup
}
