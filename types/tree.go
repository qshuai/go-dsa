package types

import (
	"container/list"
	"fmt"
	"golang.org/x/exp/constraints"
)

// Node represents tree node
type Node struct {
	Value any
	Left  *Node
	Right *Node
}

// PreOrderTraverse 前序遍历
func PreOrderTraverse(tree *Node) {
	if tree == nil {
		return
	}

	fmt.Println(tree.Value)
	PreOrderTraverse(tree.Left)
	PreOrderTraverse(tree.Right)
}

// InOrderTraverse 中序遍历
func InOrderTraverse(tree *Node) {
	if tree == nil {
		return
	}

	InOrderTraverse(tree.Left)
	fmt.Println(tree.Value)
	InOrderTraverse(tree.Right)
}

// PostOrderTraverse 后序遍历
func PostOrderTraverse(tree *Node) {
	if tree == nil {
		return
	}

	PostOrderTraverse(tree.Left)
	PostOrderTraverse(tree.Right)
	fmt.Println(tree.Value)
}

// BreadthFirstSearch 广度优先搜索BSF
func BreadthFirstSearch(tree *Node) {
	if tree == nil {
		return
	}

	l := new(list.List)
	l.PushBack(tree)

	for l.Len() > 0 {
		ele := l.Front().Value.(*Node)
		l.Remove(l.Front())

		if ele.Left != nil {
			l.PushBack(ele.Left)
		}
		if ele.Right != nil {
			l.PushBack(ele.Right)
		}

		fmt.Println(ele.Value)
	}
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
