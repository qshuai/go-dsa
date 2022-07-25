package leetcode

import (
	"container/list"

	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

// Given the root of a binary tree, return the zigzag level order traversal of its nodes'
// values. (i.e., from left to right, then right to left for the next level and alternate between).
//
// Constraints:
// The number of nodes in the tree is in the range [0, 2000].
// -100 <= TreeNode.val <= 100
func zigzagLevelOrder(root *types.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	eles := list.New()
	eles.PushBack(root)
	ans := make([][]int, 0, 1)
	var direction bool
	for eles.Len() > 0 {
		length := eles.Len()
		values := make([]int, length)
		for i := 0; i < length; i++ {
			top := eles.Front()
			eles.Remove(top)

			node := top.Value.(*types.TreeNode)
			if node.Left != nil {
				eles.PushBack(node.Left)
			}
			if node.Right != nil {
				eles.PushBack(node.Right)
			}

			if direction {
				values[length-i-1] = node.Value.(int)
			} else {
				values[i] = node.Value.(int)
			}
		}

		ans = append(ans, values)
		direction = !direction
	}

	return ans
}
