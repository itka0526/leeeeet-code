package tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	dfs := func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		return isSameTree(node1.Left, node2.Left) && isSameTree(node1.Right, node2.Right)
	}

	return dfs(p, q)
}
