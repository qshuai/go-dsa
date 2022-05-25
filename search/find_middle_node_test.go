package search

import (
	"fmt"

	"github.com/qshuai/go-dsa/types"
)

func ExampleFindMiddleNode() {
	node := types.NewListNode(1, 5)
	fmt.Println(FindMiddleNode(node).Value)

	// Output:
	// 3
}
