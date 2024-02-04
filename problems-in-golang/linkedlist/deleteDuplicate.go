package linkedlist

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	m := map[int]int{}
	curr := head
	for curr != nil {
		m[curr.Val]++
		curr = curr.Next
	}
	for k, v := range m {
		count := v
		for count > 1 {
			c := head
			for c != nil {
				if c.Val == k {
					// Unlink
					c.Next = c.Next.Next
					break
				}
				c = c.Next
			}
			count--
		}
	}
	return head
}

func TestDeleteDuplicates() {
	nums := []int{1, 1, 1}
	var head *ListNode

	i := 0
	for i < len(nums) {
		if i == 0 {
			head = &ListNode{Val: nums[i], Next: nil}
			i++
			continue
		}
		n := &ListNode{Val: nums[i]}

		c := head
		for c.Next != nil {
			c = c.Next
		}

		c.Next = n
		i++
	}

	fmt.Print("Input: \t\t")
	c := head
	for c != nil {
		fmt.Print(c.Val)
		c = c.Next
	}
	fmt.Print("\n")
	fmt.Print("Output: \t")
	deleteDuplicates(head)
	c = head
	for c != nil {
		fmt.Print(c.Val)
		c = c.Next
	}
	fmt.Print("\n")
	fmt.Println("Expected:  [1]")
}
