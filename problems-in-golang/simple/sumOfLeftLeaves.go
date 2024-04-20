package simple

func dfsLeftSum(node *TreeNode, isLeft bool, ans *int) {
	if node == nil {
		return
	}
	if isLeft && node.Left == nil && node.Right == nil {
		*ans += node.Val
	}
	dfsLeftSum(node.Left, true, ans)
	dfsLeftSum(node.Right, false, ans)
}

func sumOfLeftLeaves(root *TreeNode) (ans int) {
	dfsLeftSum(root, false, &ans)
	return ans
}
