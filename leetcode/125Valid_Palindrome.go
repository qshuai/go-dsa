package leetcode

import (
	"reflect"
	"unsafe"

	"github.com/qshuai/go-dsa/types"
)

const diffAlpha byte = 'a' - 'A'

// isPalindrome 判断字符串是否为回文串
// https://leetcode.com/problems/valid-palindrome/
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	l, size, middleNode := builDListNodes(s)
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

// builDListNodes 构建双向链表，并返回基本信息
// returns:
// 1. 链表头结点
// 2. 元素数量size
// 3. 中间节点元素(位置为 size / 2，结果四舍五入)
func builDListNodes(s string) (*types.DListNode, int, *types.DListNode) {
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

// isAlphanumeric 判断直接类型
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
