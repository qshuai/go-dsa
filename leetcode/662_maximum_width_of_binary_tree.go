package leetcode

import (
	"container/list"

	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/maximum-width-of-binary-tree/

// widthOfBinaryTree Given the root of a binary tree, return the maximum width of the given tree.
// The maximum width of a tree is the maximum width among all levels.
// The width of one level is defined as the length between the end-nodes (the leftmost
// and rightmost non-null nodes), where the null nodes between the end-nodes that would
// be present in a complete binary tree extending down to that level are also counted
// into the length calculation.
// It is guaranteed that the answer will in the range of a 32-bit signed integer.
//
// Constraints:
// The number of nodes in the tree is in the range [1, 3000].
// -100 <= Node.val <= 100
func widthOfBinaryTree(root *types.Node) int {
	type NodeWithIndex struct {
		*types.Node
		idx int
	}

	eles := list.New()
	eles.PushBack(&NodeWithIndex{
		Node: root,
		idx:  0,
	})
	var res int
	for eles.Len() > 0 {
		start := eles.Front().Value.(*NodeWithIndex).idx
		end := eles.Back().Value.(*NodeWithIndex).idx
		cnt := end - start + 1
		if cnt > res {
			res = cnt
		}

		length := eles.Len()
		for i := 0; i < length; i++ {
			item := eles.Front()
			eles.Remove(item)

			node := item.Value.(*NodeWithIndex)
			if node.Left != nil {
				eles.PushBack(&NodeWithIndex{node.Left, node.idx<<1 + 1})
			}
			if node.Right != nil {
				eles.PushBack(&NodeWithIndex{node.Right, node.idx<<1 + 2})
			}
		}
	}

	return res
}
