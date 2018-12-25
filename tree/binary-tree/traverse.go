package binary_tree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func PreOrderTraverse(tree *Node) {
	if tree == nil {
		return
	}

	fmt.Println(tree.Value)
	PreOrderTraverse(tree.Left)
	PreOrderTraverse(tree.Right)
}

func InOrderTraverse(tree *Node) {
	if tree == nil {
		return
	}

	InOrderTraverse(tree.Left)
	fmt.Println(tree.Value)
	InOrderTraverse(tree.Right)
}

func PostOrderTraverse(tree *Node) {
	if tree == nil {
		return
	}

	PostOrderTraverse(tree.Left)
	PostOrderTraverse(tree.Right)
	fmt.Println(tree.Value)
}
