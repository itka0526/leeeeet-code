package linkedlist

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 2 -> 4 -> 3
	// 5 -> 6 -> 4
	// 7 -> 0 -> 8

	// 	Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	//              +9,9,9,9
	// Output: 		[8,9,9,9,0,0,0,1]
	h := &ListNode{}
	tmp := h

	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}
	return h
}

func TestAddTwoNumbers() {
	s1 := []int{9, 9, 9, 9, 9, 9, 9}
	s2 := []int{9, 9, 9, 9}

	h1 := FromSliceToLinkedList(s1)
	h2 := FromSliceToLinkedList(s2)

	addTwoNumbers(h1, h2)
}

type Head = *ListNode

func FromSliceToLinkedList(s []int) Head {
	tmp := make([]*ListNode, len(s))
	for i := 0; i < len(s); i++ {
		tmp[i] = &ListNode{Val: s[i]}
	}

	// After storing all the nodes in a slice, we must link them
	for i, n := range tmp {
		if i+1 > len(tmp)-1 {
			continue
		}
		n.Next = tmp[i+1]
	}

	head := tmp[0]
	return head
}

func TraverseLinkedList(n *ListNode) {
	curr := n
	fmt.Printf("----------- Linked List Start: \t%p -----------\n", (n))
	for curr != nil {
		fmt.Println(curr)
		curr = curr.Next
	}
	fmt.Printf("----------- Linked List End: \t%p -----------\n", (n))
}
