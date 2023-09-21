package tree

func diameterOfBinaryTree(root *TreeNode) int {
	var maxLength int
	var dfs func(*TreeNode, *int) int

	dfs = func(tn *TreeNode, length *int) int {
		if tn == nil {
			return 0
		}
		left := dfs(tn.Left, length)
		right := dfs(tn.Right, length)
		if *length < left+right {
			*length = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}

	dfs(root, &maxLength)

	return maxLength
}
