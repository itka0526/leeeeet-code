package tree

import (
	"fmt"
	"log"
	"sort"
)

func kthSmallest(root *TreeNode, k int) int {
	// add everything to a slice and sort it then get the k-th element
	// I feel like I could have somehow used BST
	// l := make([]*TreeNode, 0)

	// var dfs func(n *TreeNode)
	// dfs = func(n *TreeNode) {
	// 	if n == nil {
	// 		return
	// 	}
	// 	dfs(n.Left)
	// 	dfs(n.Right)
	// 	l = append(l, n)
	// }

	// dfs(root)

	// sort.Slice(l, func(i, j int) bool {
	// 	return l[i].Val < l[j].Val
	// })

	// return l[k].Val

	// stack := []*TreeNode{}
	// curr := root

	// // traverse the left subtree and once you reach a null pointer go back
	// // add that element to the stack
	// for {
	// 	for curr != nil {
	// 		stack = append(stack, curr)
	// 		curr = curr.Left
	// 	}

	// 	curr = stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]

	// 	k--
	// 	if k == 0 {
	// 		return curr.Val
	// 	}

	// 	curr = curr.Right
	// }

	// neetcode's solution is weird it uses iterative solution I am not sure how its working but seems complicated

	// I just wanna use the fact that we have a Binary Tree
	s := make([]*TreeNode, 0)
	addElementsInOrder(root, &s)
	return s[k-1].Val
}

func addElementsInOrder(n *TreeNode, s *[]*TreeNode) {
	if n.Left != nil {
		addElementsInOrder(n.Left, s)
	}
	*s = append(*s, n)
	if n.Right != nil {
		addElementsInOrder(n.Right, s)
	}
}

func TestKthSmallestElementInBST() {
	s := []int{3, 1, 4, 2}
	k := 1

	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	log.Println("Initial sorted slice: ", s)

	root := CreateBST(s, 0, len(s)-2)
	res := kthSmallest(root, k)
	fmt.Println(res)
}
