package list

import (
	"fmt"

	"github.com/qshuai/go-dsa/types"
)

func ExampleFindMiddleNode() {
	node := &types.ListNode{Value: 1, Next: &types.ListNode{Value: 2, Next: &types.ListNode{Value: 3, Next: &types.ListNode{Value: 4, Next: &types.ListNode{Value: 5, Next: nil}}}}}
	fmt.Println(findMiddleNode(node).Value)

	// Output:
	// 3
}
