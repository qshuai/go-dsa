package leetcode

import (
	"reflect"
	"unsafe"

	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/valid-palindrome/

const diffAlpha byte = 'a' - 'A'

// isPalindrome 判断字符串是否为回文串
// solution 1：使用头尾指针实现
// Best practice
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	i, j := 0, len(s)-1
	for j > i {
		// 从前往后寻找alphanumeric
		var a byte
		for i <= j {
			ok, isUpper := isAlphanumeric(s[i])
			if !ok {
				i++
				continue
			}

			if isUpper {
				a = s[i] + diffAlpha
			} else {
				a = s[i]
			}
			break
		}

		// 从后往前寻找alphanumeric
		var b byte
		for i <= j {
			ok, isUpper := isAlphanumeric(s[j])
			if !ok {
				j--
				continue
			}

			if isUpper {
				b = s[j] + diffAlpha
			} else {
				b = s[j]
			}
			break
		}

		if a != b {
			return false
		}

		i++
		j--
	}

	return true
}

// solution 2: 使用链表实现
func isPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}

	l, size, middleNode := buildListNodes(s)
	if l == nil {
		return true
	}

	if size%2 == 0 {
		// 偶数个元素
		prev := middleNode
		next := middleNode.Next

		for prev != nil {
			if prev.Value != next.Value {
				return false
			}

			prev = prev.Prev
			next = next.Next
		}
	} else {
		// 奇数个元素
		prev := middleNode.Prev
		next := middleNode.Next

		for prev != nil {
			if prev.Value != next.Value {
				return false
			}

			prev = prev.Prev
			next = next.Next
		}
	}

	return true
}

// buildListNodes 构建双向链表，并返回基本信息
// returns:
// 1. 链表头结点
// 2. 元素数量size
// 3. 中间节点元素(位置为 size / 2，结果四舍五入)
func buildListNodes(s string) (*types.DListNode, int, *types.DListNode) {
	var l *types.DListNode

	// 将字符串转换成byte数组
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceHeader := &reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	bs := *(*[]byte)(unsafe.Pointer(sliceHeader))

	var currentNode *types.DListNode
	var middleNode *types.DListNode
	var size int
	for _, b := range bs {
		ok, isUpper := isAlphanumeric(b)
		if !ok {
			continue
		}

		size++
		if isUpper {
			// 大写字母转小写
			b = b + diffAlpha
		}

		if currentNode == nil {
			l = &types.DListNode{
				Value: b,
			}

			currentNode = l
			middleNode = l
			continue
		}

		newNode := &types.DListNode{
			Value: b,
			Prev:  currentNode,
		}

		currentNode.Next = newNode
		currentNode = newNode

		// 维护中间节点
		if size%2 == 1 {
			middleNode = middleNode.Next
		}
	}

	return l, size, middleNode
}

// isAlphanumeric 判断字节类型
// return:
// 1. 是否为字母数字类型的直接
// 2. 判断是否为大写字母
func isAlphanumeric(b byte) (bool, bool) {
	if (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') {
		return true, false
	}

	if b >= 'A' && b <= 'Z' {
		return true, true
	}

	return false, false
}

// isAlphanumericRune 判断字符类型
// return:
// 1. 是否为字母数字类型的直接
// 2. 判断是否为大写字母
func isAlphanumericRune(b rune) (bool, bool) {
	if (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') {
		return true, false
	}

	if b >= 'A' && b <= 'Z' {
		return true, true
	}

	return false, false
}
