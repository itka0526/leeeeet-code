package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// two pointer technique comes to the rescue
	if head == nil {
		return head
	}

	left, right := head, head

	// increase distance between the two pointers by N
	for n > 0 {
		right = right.Next
		n--
	}

	if right == nil {
		return left.Next
	}

	// till right pointer reaches null iterate and if it reaches that means left pointer is at the correct position
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	// now we just reconnect the left to the next next node, left just jumps over one node.
	left.Next = left.Next.Next
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	if head == nil {
		return head
	}

	for n > 0 {
		fast = fast.Next
		n--
	}

	if fast == nil {
		return slow.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}
