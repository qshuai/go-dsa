package leetcode

// https://leetcode.com/problems/generate-parentheses/

const openParenthesis byte = '('
const closeParenthesis byte = ')'

// generateParenthesis Given n pairs of parentheses, write a function to generate
// all combinations of well-formed parentheses.
// Constraints:
// 1 <= n <= 8
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	generate(n, 0, 0, "", &res)

	return res
}

func generate(n int, opens, closes int, target string, collector *[]string) {
	if opens == closes {
		if opens == n {
			*collector = append(*collector, target)
			return
		}

		generate(n, opens+1, closes, target+"(", collector)
		return
	}

	if opens == n {
		generate(n, opens, closes+1, target+")", collector)
		return
	}

	generate(n, opens+1, closes, target+"(", collector)
	generate(n, opens, closes+1, target+")", collector)
}
