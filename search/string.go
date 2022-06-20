package search

import (
	"math/big"

	"github.com/qshuai/go-dsa/utils"
)

// IndexOfUsingBF returns the index of the first instance of sub string
// from str, or -1 if not matched substring present from str.
// Brute Force algorithm
func IndexOfUsingBF(str, sub string) int {
	if len(str) < len(sub) || len(sub) <= 0 {
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
	if len(sub) > len(str) || len(sub) <= 0 {
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
	for i := 1; i < len(sub); i++ {
		temp := big.NewInt(base)
		hashes[i] = temp.Mul(temp, prev)

		prev = hashes[i]
	}

	// 计算主串初始、模式串的hash值
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

	multiple := big.NewInt(base)
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
// Boyer-Moore algorithm
// Constraints:
// str and sub are consist of ascii character
func IndexOfUsingBM(str, sub string) int {
	if len(str) < len(sub) || len(sub) <= 0 {
		return -1
	}

	bc := buildAsciiPos(sub)
	suffix, prefix := buildMatchedSegment(sub)
	for i := 0; i <= len(str)-len(sub); {
		j := len(sub) - 1
		// 将模式串按照从后向前的顺序和主串进行比对，如果出现不匹配的情况，此时j就代表坏字符在模式串中的索引位置
		for ; j >= 0; j-- {
			if sub[j] != str[i+j] {
				break
			}
		}

		if j < 0 {
			// 索引上面的for循环没有出现break的情况，此时j为-1
			return i
		}

		// 按照坏字符规则，需要移动的步数
		x := j - bc[str[i+j]] + 1

		var y int
		// j < len(sub)-1代表上面匹配比较时，出现模式串的后缀子串和主串有部分匹配
		if j < len(sub)-1 {
			y = moveByMatchedSegment(j, len(sub), suffix, prefix)
		}

		// 取坏字符规则、好后缀规则的最大移动步数
		i += utils.Max(x, y)
		utils.Max("", "")
	}

	return -1
}

// IndexOfUsingKMP  returns the index of the first instance of sub string
// from str, or -1 if not matched substring present from str.
// KMP algorithm(Knuth, Morris and Pratt)
func IndexOfUsingKMP(str, sub string) int {
	if len(str) < len(sub) || len(sub) <= 0 {
		return -1
	}

	next := buildNextArray(sub)
	var j int
	for i := 0; i < len(str); {
		if str[i] == sub[j] {
			i++
			j++
		} else if j > 0 {
			j = next[j-1]
		} else {
			i++
		}

		if j == len(sub) {
			return i - j
		}
	}

	return -1
}

// buildAsciiArray 构建字符串的ascii码值和位置的hash表。数组的索引代表字符的ascii码值，值为
// 该字符在字符串中最后的索引+1（+1的目的是使数组的默认值0能够代表字符不存在的含义）
func buildAsciiPos(m string) []int {
	if len(m) <= 0 {
		return nil
	}

	ret := make([]int, 256)
	for i := 0; i < len(m); i++ {
		ret[m[i]] = i + 1
	}

	return ret
}

// buildMatchedSegment 在模式串中构建和后缀子串匹配的前向字符串位置信息。suffix数组：索引代表匹配的长度，
// 值为位置匹配的前向字符串起始索引+1（+1的目的是是数组的默认值0能够代表该长度的后缀子串没有匹配的前向字符串）。
// prefix数组：索引代表匹配的长度，值如果为true，代表该匹配长度可以和模式串的前缀子串重合。
func buildMatchedSegment(m string) ([]int, []bool) {
	suffix := make([]int, len(m))
	prefix := make([]bool, len(m))
	for i := 0; i < len(m)-1; i++ {
		// k代表匹配串的长度
		j, k := i, 0
		for j >= 0 && m[j] == m[len(m)-1-k] {
			suffix[k+1] = j + 1

			j--
			k++
		}

		if j == -1 {
			prefix[k] = true
		}
	}

	return suffix, prefix
}

// moveByMatchedSegment 按照好后缀的规则计算向后移动的步数
func moveByMatchedSegment(j, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j
	if suffix[k] != 0 {
		// 代表好后缀存在匹配的前向子字符串
		return j - (suffix[k] - 1) + 1
	}

	// 如果整个好后缀没有可匹配的前向子字符串的话，只能将好后缀的后缀子串和模式串的前缀子串进行尝试匹配
	for i := j + 2; i < m-1; i++ {
		if prefix[m-i] {
			return i
		}
	}

	return m
}

// buildNextArray 获取KMP算法的next数组。索引和模式串的索引保持一致，
// 值为该索引对应的模式串的前缀子串中最长公共前后缀子串长度。例如索引为i，对应
// 的模式串子串为m[0:i]，如果该子串为"ababa"，那么最长匹配的公共前后缀
// 长度为3（其中前缀为aba，后缀也为aba）
func buildNextArray(m string) []int {
	if len(m) <= 0 {
		return nil
	}

	next := make([]int, 0, len(m))
	next = append(next, 0)
	var k int
	for i := 1; i < len(m); {
		if m[k] == m[i] {
			// 如果匹配，将长度+1，游标向后移动一位
			k++
			i++
			next = append(next, k)
		} else {
			if k == 0 {
				// 之前没有可匹配的了，将当前匹配长度设置为0，并将游标向下移动一位
				next = append(next, 0)
				i++
			} else {
				// 使用前一位的匹配结果，继续检查
				k = next[k-1]
			}
		}
	}

	return next
}
