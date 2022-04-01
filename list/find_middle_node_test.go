package list

import (
	"fmt"

	"github.com/qshuai/go-dsa/types"
)

func ExampleFindMiddleNode() {
	node := types.NewListNode(1, 5)
	fmt.Println(findMiddleNode(node).Value)

	// Output:
	// 3
}
