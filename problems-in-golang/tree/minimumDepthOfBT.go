package tree

import "math"

func minDepth(root *TreeNode) int {
	min := math.MaxInt
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, c int) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil && min > c {
			min = c
		}
		dfs(n.Left, c+1)
		dfs(n.Right, c+1)
	}
	dfs(root, 1)
	return min
}
