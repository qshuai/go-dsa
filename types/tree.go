package types

import (
	"container/list"
	"fmt"

	"golang.org/x/exp/constraints"
)

// TreeNode represents tree node
type TreeNode struct {
	Val   any
	Left  *TreeNode
	Right *TreeNode
}

// PreOrderTraverse 前序遍历
func PreOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Println(tree.Val)
	PreOrderTraverse(tree.Left)
	PreOrderTraverse(tree.Right)
}

// PreOrderLoopTraverse 使用栈的方式进行前序遍历
func PreOrderLoopTraverse(tree *TreeNode) {
	container := list.New()
	node := tree
	for node != nil || container.Len() > 0 {
		for node != nil {
			fmt.Println(node.Val)

			container.PushBack(node)
			node = node.Left
		}

		if container.Len() > 0 {
			ele := container.Back()
			node = ele.Value.(*TreeNode)
			container.Remove(ele)

			node = node.Right
		}
	}
}

// InOrderTraverse 中序遍历
func InOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}

	InOrderTraverse(tree.Left)
	fmt.Println(tree.Val)
	InOrderTraverse(tree.Right)
}

// InOrderLoopTraverse 使用栈的方式进行中序遍历
func InOrderLoopTraverse(tree *TreeNode) {
	container := list.New()
	node := tree
	for node != nil || container.Len() > 0 {
		for node != nil {
			container.PushBack(node)
			node = node.Left
		}

		if container.Len() > 0 {
			ele := container.Back()
			node = ele.Value.(*TreeNode)
			container.Remove(ele)

			fmt.Println(node.Val)

			node = node.Right
		}
	}
}

// PostOrderTraverse 后序遍历
func PostOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}

	PostOrderTraverse(tree.Left)
	PostOrderTraverse(tree.Right)
	fmt.Println(tree.Val)
}

// PostOrderLoopTraverse 使用栈的方式进行后序遍历
func PostOrderLoopTraverse(tree *TreeNode) {
	container := list.New()
	node := tree
	lastVisit := tree
	for node != nil || container.Len() > 0 {
		for node != nil {
			container.PushBack(node)
			node = node.Left
		}

		ele := container.Back()
		node = ele.Value.(*TreeNode)
		if node.Right == nil || node.Right == lastVisit {
			fmt.Println(node.Val)
			container.Remove(ele)

			lastVisit = node
			node = nil
		} else {
			node = node.Right
		}
	}
}

// BreadthFirstSearch 广度优先搜索BSF
func BreadthFirstSearch(tree *TreeNode) {
	if tree == nil {
		return
	}

	l := new(list.List)
	l.PushBack(tree)

	for l.Len() > 0 {
		ele := l.Front().Value.(*TreeNode)
		l.Remove(l.Front())

		if ele.Left != nil {
			l.PushBack(ele.Left)
		}
		if ele.Right != nil {
			l.PushBack(ele.Right)
		}

		fmt.Println(ele.Val)
	}
}

// NewBinaryTree 根据数组生成二叉树，传入数组的元素顺序需要和广度优先遍历的顺序一致
func NewBinaryTree(vals []any) *TreeNode {
	if len(vals) <= 0 || vals[0] == nil {
		return nil
	}

	head := &TreeNode{Val: vals[0]}
	queue := list.New()
	queue.PushBack(head)
	i := 1
	for queue.Len() > 0 && i < len(vals) {
		item := queue.Front()
		queue.Remove(item)
		ele := item.Value.(*TreeNode)
		if vals[i] != nil {
			ele.Left = &TreeNode{Val: vals[i]}
		}
		queue.PushBack(ele.Left)

		if i+1 < len(vals) && vals[i+1] != nil {
			ele.Right = &TreeNode{Val: vals[i+1]}
		}
		queue.PushBack(ele.Right)

		i += 2
	}

	return head
}

// BST represents binary search tree
type BST[T constraints.Ordered] struct {
	Value T
	Left  *BST[T]
	Right *BST[T]
}

// Insert pushes a new element to bst
func (bst *BST[T]) Insert(value T) {
	if bst == nil {
		panic("insert into nil BST not supported")
	}

	cursor := bst
	for cursor != nil {
		if value < cursor.Value {
			if cursor.Left == nil {
				cursor.Left = &BST[T]{Value: value}
				return
			}
			cursor = cursor.Left
			continue
		}

		if cursor.Right == nil {
			cursor.Right = &BST[T]{Value: value}
			return
		}
		cursor = cursor.Right
	}
}

// Find return the tree node that is equal to the target, nil returned if not found
func (bst *BST[T]) Find(target T) *BST[T] {
	if bst == nil {
		return nil
	}

	if bst.Value == target {
		return bst
	}
	if bst.Value < target {
		if bst.Left == nil {
			return nil
		}

		return bst.Left.Find(target)
	}

	if bst.Right == nil {
		return nil
	}

	return bst.Right.Find(target)
}

// Height returns the height of tree(0-based)
func (bst *BST[T]) Height() int {
	if bst == nil {
		return 0
	}

	var max int
	bst.height(&max, -1)
	return max
}

// BuildBSTFromSlice builds a binary search tree from slice
func BuildBSTFromSlice[T constraints.Ordered](s []T) *BST[T] {
	if len(s) <= 0 {
		return nil
	}

	root := &BST[T]{Value: s[0]}
	for i := 1; i < len(s); i++ {
		root.Insert(s[i])
	}

	return root
}

func (bst *BST[T]) height(max *int, height int) {
	if bst == nil {
		if height > *max {
			*max = height
		}

		return
	}

	bst.Left.height(max, height+1)
	bst.Right.height(max, height+1)
}
