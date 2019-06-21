package list

import (
	"fmt"
)

func ExampleFindMiddleNode() {
	node := &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: nil}}}}}
	fmt.Println(findMiddleNode(node).Value)

	// Output:
	// 3
}
