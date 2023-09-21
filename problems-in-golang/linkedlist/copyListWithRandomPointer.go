package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	unlinkedNodes := make(map[*Node]*Node)
	curr := head

	// copied nodes will be assigned to the existing nodes
	for curr != nil {
		// I do not understand this section
		copy := &Node{Val: curr.Val}
		unlinkedNodes[curr] = copy
		curr = curr.Next
	}

	// reseting current to head
	curr = head

	for curr != nil {
		// now I must link them all if I did not misunderstand
		copiedNode := unlinkedNodes[curr]
		copiedNode.Next = unlinkedNodes[curr.Next]
		copiedNode.Random = unlinkedNodes[curr.Random]
		curr = curr.Next
	}

	return head
}

func copyRandomList1(head *Node) *Node {
	copiedNodes := make(map[*Node]*Node)
	curr := head

	for curr != nil {
		newNode := Node{Val: curr.Val}
		copiedNodes[curr] = &newNode
		curr = curr.Next
	}

	curr = head

	for curr != nil {
		copiedNode := copiedNodes[curr]
		copiedNode.Next = copiedNodes[curr.Next]
		copiedNode.Random = copiedNodes[curr.Random]
		curr = curr.Next
	}

	return copiedNodes[head]
}
