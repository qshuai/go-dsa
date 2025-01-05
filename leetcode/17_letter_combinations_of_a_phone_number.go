package leetcode

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

var letterMapping = [10][]string{nil, nil, {"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"},
	{"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}

const baseDigit byte = '0'

// letterCombinations Given a string containing digits from 2-9 inclusive, return all possible letter combinations
// that the number could represent. Return the answer in any order.
// A mapping of digit to letters (just like on the telephone buttons) is given below.
// Note that 1 does not map to any letters.
func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}

	res := make([]string, 0)
	combine(digits, 0, "", &res)

	return res
}

func combine(data string, idx int, middle string, collector *[]string) {
	if len(middle) >= len(data) {
		*collector = append(*collector, middle)
		return
	}

	bs := letterMapping[data[idx]-baseDigit]
	for i := 0; i < len(bs); i++ {
		combine(data, idx+1, middle+bs[i], collector)
	}
}
