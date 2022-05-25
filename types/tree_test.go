package types

var root = buildTree()

func ExamplePreOrderTraverse() {
	PreOrderTraverse(root)
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
	InOrderTraverse(root)
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
	PostOrderTraverse(root)
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
	BreadthFirstSearch(root)
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
	root := Node{1, nil, nil}
	root.Left = &Node{2, nil, nil}
	root.Right = &Node{3, nil, nil}

	root.Left.Left = &Node{4, nil, nil}
	root.Left.Right = &Node{5, nil, nil}
	root.Left.Right.Left = &Node{8, nil, nil}

	root.Right.Left = &Node{6, nil, nil}
	root.Right.Right = &Node{7, nil, nil}
	root.Right.Right.Right = &Node{9, nil, nil}

	return &root
}
