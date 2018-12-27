package binary_tree

func buildTree() *Node {
	// Illustration：
	//                  1
	//              /      \
	//             2        3
	//           /   \    /   \
	//          4     5  6     7
	//               /           \
	//              8             9
	//
	tree := Node{1, nil, nil}
	tree.Left = &Node{2, nil, nil}
	tree.Right = &Node{3, nil, nil}

	tree.Left.Left = &Node{4, nil, nil}
	tree.Left.Right = &Node{5, nil, nil}
	tree.Left.Right.Left = &Node{8, nil, nil}

	tree.Right.Left = &Node{6, nil, nil}
	tree.Right.Right = &Node{7, nil, nil}
	tree.Right.Right.Right = &Node{9, nil, nil}

	return &tree
}

func ExamplePreOrderTraverse() {
	tree := buildTree()

	PreOrderTraverse(tree)
	// Output:
	// 1
	// 2
	// 4
	// 5
	// 8
	// 3
	// 6
	// 7
	// 9
}

func ExampleInOrderTraverse() {
	tree := buildTree()

	InOrderTraverse(tree)
	// Output:
	// 4
	// 2
	// 8
	// 5
	// 1
	// 6
	// 3
	// 7
	// 9
}

func ExamplePostOrderTraverse() {
	tree := buildTree()

	PostOrderTraverse(tree)
	// Output:
	// 4
	// 8
	// 5
	// 2
	// 6
	// 9
	// 7
	// 3
	// 1
}

func ExampleLayerOrderTraverse() {
	tree := buildTree()

	BreadthFirstSearch(tree)
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}
