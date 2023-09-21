package linkedlist

import "github.com/itka0526/problems-in-golang/functions"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
}

func (l *List) add(value int) {
	newNode := &ListNode{Val: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

func reverseList(head *ListNode) *ListNode {
	return head
}

func createLL() *List {
	slice := []int{1, 2, 3, 4, 5}
	list := &List{}

	for _, val := range slice {
		list.add(val)
	}

	return list
}

func TestReverseList() {
	functions.PrintSolution(reverseList(createLL().head))
}
