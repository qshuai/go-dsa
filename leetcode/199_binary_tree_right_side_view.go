package leetcode

import (
	"container/list"

	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/binary-tree-right-side-view/

// rightSideView Given the root of a binary tree, imagine yourself standing on the right side of it,
// return the values of the nodes you can see ordered from top to bottom.
//
// Constraints:
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
func rightSideView(root *types.TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)
	ans := make([]int, 0, 1)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			ele := queue.Front()
			queue.Remove(ele)
			item := ele.Value.(*types.TreeNode)
			if item.Left != nil {
				queue.PushBack(item.Left)
			}
			if item.Right != nil {
				queue.PushBack(item.Right)
			}

			// 最后一个元素
			if i == length-1 {
				ans = append(ans, item.Val.(int))
			}
		}
	}

	return ans
}
