package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/sum-root-to-leaf-numbers/

// sumNumbers You are given the root of a binary tree containing digits from 0 to 9 only.
// Each root-to-leaf path in the tree represents a number.
// For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
// Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.
// A leaf node is a node with no children.
//
// Constraints:
// The number of nodes in the tree is in the range [1, 1000].
// 0 <= Node.val <= 9
// The depth of the tree will not exceed 10.
func sumNumbers(root *types.TreeNode) int {
	var sum int
	dfsPathDistance(root, new([]int), &sum)
	return sum
}

func dfsPathDistance(node *types.TreeNode, path *[]int, sum *int) {
	*path = append(*path, node.Value.(int))
	if node.Left == nil && node.Right == nil {
		// 叶子结点
		pow := 1
		for i := len(*path) - 1; i >= 0; i-- {
			*sum += (*path)[i] * pow
			pow *= 10
		}

		return
	}

	// 左节点
	if node.Left != nil {
		dfsPathDistance(node.Left, path, sum)
		*path = (*path)[:len(*path)-1]
	}

	// 右节点
	if node.Right != nil {
		dfsPathDistance(node.Right, path, sum)
		*path = (*path)[:len(*path)-1]
	}
}
