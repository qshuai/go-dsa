package leetcode

import (
	"github.com/qshuai/go-dsa/types"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

// lowestCommonAncestor Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and
// q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
//
// Constraints:
// The number of nodes in the tree is in the range [2, 105].
// -10^9 <= Node.val <= 10^9
// All Node.val are unique.
// p != q
// p and q will exist in the tree.
func lowestCommonAncestor(root, p, q *types.TreeNode) *types.TreeNode {
	target := new(types.TreeNode)
	dfsAncestor(root, p, q, target)
	return target
}

func dfsAncestor(root, p, q, target *types.TreeNode) bool {
	if root == nil {
		return false
	}

	lson := dfsAncestor(root.Left, p, q, target)
	rson := dfsAncestor(root.Right, p, q, target)
	if (lson && rson) || ((lson || rson) && (root.Val.(int) == p.Val.(int) || root.Val.(int) == q.Val.(int))) {
		*target = *root
	}

	return lson || rson || root.Val.(int) == p.Val.(int) || root.Val.(int) == q.Val.(int)
}

// lowestCommonAncestor2 Solution-2 寻找p、q的父节点路径，通过两个路径并借助hash表得出最近公共祖先
func lowestCommonAncestor2(root, p, q *types.TreeNode) *types.TreeNode {
	temp := []*types.TreeNode{root}
	pPath, qPath := new([]*types.TreeNode), new([]*types.TreeNode)
	dfsTreePath(root, p.Val.(int), q.Val.(int), &temp, pPath, qPath)
	m := make(map[*types.TreeNode]struct{})
	for _, item := range *pPath {
		m[item] = struct{}{}
	}

	for i := len(*qPath) - 1; i >= 0; i-- {
		if _, ok := m[(*qPath)[i]]; ok {
			return (*qPath)[i]
		}
	}

	return nil
}

func dfsTreePath(node *types.TreeNode, p, q int, temp, pPath, qPath *[]*types.TreeNode) {
	if len(*pPath) > 0 && len(*qPath) > 0 {
		return
	}
	if node.Val == p {
		*pPath = append(*pPath, *temp...)
	}
	if node.Val == q {
		*qPath = append(*qPath, *temp...)
	}

	// 	遍历左节点
	if node.Left != nil {
		*temp = append(*temp, node.Left)
		dfsTreePath(node.Left, p, q, temp, pPath, qPath)
		*temp = (*temp)[:len(*temp)-1]
	}
	if node.Right != nil {
		*temp = append(*temp, node.Right)
		dfsTreePath(node.Right, p, q, temp, pPath, qPath)
		*temp = (*temp)[:len(*temp)-1]
	}
}
