package search

import (
	"math/big"

	"github.com/qshuai/go-dsa/utils"
)

// IndexOfUsingBF returns the index of the first instance of sub string
// from str, or -1 if not matched substring present from str.
// Brute Force algorithm
func IndexOfUsingBF(str, sub string) int {
	if len(str) < len(sub) {
		return -1
	}

	for i := 0; i <= len(str)-len(sub); i++ {
		matched := true
		for j := 0; j < len(sub); j++ {
			if str[i+j] != sub[j] {
				matched = false
				break
			}
		}

		if matched {
			return i
		}
	}

	return -1
}

// IndexOfUsingRK returns the index of the first instance of sub string
// from str, or -1 if not matched substring present from str.
// Rabin-Karp algorithm
func IndexOfUsingRK(str, sub string) int {
	if len(sub) > len(str) {
		return -1
	}
	if str == sub {
		return 0
	}

	// 计算26的n次方，把结果缓存起来，为了之后重复、高效的计算字符串hash值
	hashes := make([]*big.Int, len(sub))
	hashes[0] = big.NewInt(1)
	prev := hashes[0]
	const base = 26
	multiple := big.NewInt(base)
	for i := 1; i < len(sub); i++ {
		temp := big.NewInt(base)
		hashes[i] = temp.Mul(temp, prev)

		prev = hashes[i]
	}

	// 计算主串初始的hash值
	const startVal = 'a'
	modeHash := big.NewInt(0)
	hash := big.NewInt(0)
	for i := 0; i < len(sub); i++ {
		modVal := big.NewInt(int64(sub[i] - startVal))
		modeHash.Add(modeHash, modVal.Mul(modVal, hashes[i]))

		val := big.NewInt(int64(str[i] - startVal))
		hash.Add(hash, val.Mul(val, hashes[i]))
	}
	if hash.Cmp(modeHash) == 0 {
		return 0
	}

	for i := 1; i < len(str)-len(sub)+1; i++ {
		// 通过上一个位置的hash值，推算当前字符串的hash值
		hash.Sub(hash, big.NewInt(int64(str[i-1]-startPoint)))
		hash.Div(hash, multiple)
		endVal := big.NewInt(int64(str[i+len(sub)-1] - startPoint))
		hash.Add(hash, endVal.Mul(endVal, hashes[len(hashes)-1]))

		if hash.Cmp(modeHash) == 0 {
			return i
		}
	}

	return -1
}

// IndexOfUsingBM return the first index of matched sub string, or -1
// if not found
// Using Boyer-Moore algorithm
// Constraints:
// str and sub are consist of ascii character
func IndexOfUsingBM(str, sub string) int {
	if len(str) < len(sub) {
		return -1
	}

	bc := buildAsciiArray(sub)
	suffix, prefix := buildMatchedSegment(sub)
	for i := 0; i <= len(str)-len(sub); {
		j := len(sub) - 1
		for ; j >= 0; j-- {
			if sub[j] != str[i+j] {
				break
			}
		}

		if j < 0 {
			// 找到匹配的子串
			return i
		}

		x := j - bc[str[i+j]] + 1
		var y int
		if j < len(sub)-1 {
			y = moveByMatchedSegment(j, len(sub), suffix, prefix)
		}

		i += utils.Max(x, y)
	}

	return -1
}

func buildAsciiArray(m string) []int {
	if len(m) <= 0 {
		return nil
	}

	ret := make([]int, 256)
	for i := 0; i < len(m); i++ {
		ret[m[i]] = i + 1
	}

	return ret
}

func buildMatchedSegment(m string) ([]int, []bool) {
	suffix := make([]int, len(m))
	prefix := make([]bool, len(m))
	for i := 0; i < len(m)-1; i++ {
		var k int
		j := i
		for j >= 0 && m[j] == m[len(m)-1-k] {
			j--
			k++
			suffix[k] = j + 2
		}

		if j == -1 {
			prefix[k] = true
		}
	}

	return suffix, prefix
}

func moveByMatchedSegment(j, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j
	if suffix[k] != 0 {
		return j - suffix[k] + 2
	}

	for i := j + 2; i < m-1; i++ {
		if prefix[m-i] {
			return i
		}
	}

	return m
}
