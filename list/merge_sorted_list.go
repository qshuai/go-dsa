package list

// 将两个有序的链表合并，要求合并之后的链表依然有序

func mergeSortedList(l1 *Node, l2 *Node) *Node {
	if l1 == nil || l2 == nil {
		if l1 != nil {
			return l1
		}

		if l2 != nil {
			return l2
		}

		return nil
	}

	var ret *Node
	lastNode := ret
	var head *Node
	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			ret = l1
			l1 = l1.Next
		} else {
			ret = l2
			l2 = l2.Next
		}

		// the initial state
		if head == nil {
			head = ret
			lastNode = head
		} else {
			lastNode.Next = ret
			lastNode = ret
		}
	}

	if l1 != nil {
		ret = l1
	}

	if l2 != nil {
		ret = l2
	}

	lastNode.Next = ret

	return head
}
