package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/symmetric-tree/

// isSymmetric Given the root of a binary tree, check whether it is a mirror of
// itself (i.e., symmetric around its center).
func isSymmetric(root *types.TreeNode) bool {
	return root == nil || checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(left, right *types.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return checkSymmetric(left.Left, right.Right) && checkSymmetric(left.Right, right.Left)
}
