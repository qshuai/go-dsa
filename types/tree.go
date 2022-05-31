package types

import (
	"container/list"
	"fmt"
)

type Node struct {
	Value int
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
