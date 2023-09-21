package linkedlist

func mergeTwoLists(node1 *ListNode, node2 *ListNode) *ListNode {
	res := new(ListNode)
	tail := res

	for node1 != nil && node2 != nil {
		if node2.Val >= node1.Val {
			tail.Next = node1
			node1 = node1.Next
		} else {
			tail.Next = node2
			node2 = node2.Next
		}
		tail = tail.Next
	}

	if node1 != nil {
		tail.Next = node1
	}

	if node2 != nil {
		tail.Next = node2
	}

	res = res.Next

	return res
}
