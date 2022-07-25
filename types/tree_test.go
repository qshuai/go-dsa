package types

import (
	"reflect"
	"testing"
)

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

func ExamplePreOrderLoopTraverse() {
	PreOrderLoopTraverse(root)
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

func ExampleInOrderLoopTraverse() {
	InOrderLoopTraverse(root)
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

func ExamplePostOrderLoopTraverse() {
	PostOrderLoopTraverse(root)
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

func buildTree() *TreeNode {
	// Illustration：
	//                 1
	//              /      \
	//             2        3
	//           /   \    /   \
	//          4     5  6     7
	//               /           \
	//              8             9
	//
	root := TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}

	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	root.Left.Right.Left = &TreeNode{8, nil, nil}

	root.Right.Left = &TreeNode{6, nil, nil}
	root.Right.Right = &TreeNode{7, nil, nil}
	root.Right.Right.Right = &TreeNode{9, nil, nil}

	return &root
}

func TestBuildBSTFromSlice(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want *BST[int]
	}{
		{
			name: "case-1",
			args: []int{1, 2, 3},
			want: &BST[int]{Value: 1, Right: &BST[int]{Value: 2, Right: &BST[int]{Value: 3}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildBSTFromSlice(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildBSTFromSlice() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestBST_Height(t *testing.T) {
	tests := []struct {
		name string
		args *BST[int]
		want int
	}{
		{
			name: "case-1",
			args: BuildBSTFromSlice([]int{1, 2, 3}),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.Height(); got != tt.want {
				t.Errorf("Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want *TreeNode
	}{
		{
			name: "case-1",
			args: []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
			want: &TreeNode{
				Value: 3,
				Left: &TreeNode{
					Value: 5,
					Left: &TreeNode{
						Value: 6,
					},
					Right: &TreeNode{
						Value: 2,
						Left: &TreeNode{
							Value: 7,
						},
						Right: &TreeNode{
							Value: 4,
						},
					},
				},
				Right: &TreeNode{
					Value: 1,
					Left: &TreeNode{
						Value: 0,
					},
					Right: &TreeNode{
						Value: 8,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinaryTree(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
