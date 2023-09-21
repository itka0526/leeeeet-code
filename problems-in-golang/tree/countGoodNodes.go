package tree

var dfs func(*TreeNode, int)

func goodNodes(root *TreeNode) int {
	count := 0
	dfs = func(node *TreeNode, maxVal int) {
		if node == nil {
			return
		}
		if node.Val >= maxVal {
			count++
		}
		// maxVal must be updated! I feel dumb!
		dfs(node.Left, max(maxVal, node.Val))
		dfs(node.Right, max(maxVal, node.Val))
	}
	dfs(root, root.Val)
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
