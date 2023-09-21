package linkedlist

func reorderList(head *ListNode) {
	nodes := make([]*ListNode, 0)

	curr := head

	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	left, right := 0, len(nodes)-1

	for left < right {
		nodes[right].Next = nil
		nodes[left].Next = nodes[right]
		left++
		nodes[right].Next = nodes[left]
		nodes[left].Next = nil
		right--
	}

	head = nodes[0]
}
