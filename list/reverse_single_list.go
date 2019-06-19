package list

func ReverseSingleList(list *Node) *Node {
	cur := list
	var prev *Node

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}
