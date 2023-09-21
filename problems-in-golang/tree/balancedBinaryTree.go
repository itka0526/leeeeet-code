package tree

import "math"

type BalanceTree struct {
	Balance bool
	Height  int
}

func isBalanced(root *TreeNode) bool {
	var dfs func(*TreeNode) BalanceTree

	dfs = func(n *TreeNode) BalanceTree {
		if n == nil {
			return BalanceTree{Balance: true, Height: 0}
		}
		left, right := dfs(n.Left), dfs(n.Right)
		balanced := (left.Balance && right.Balance && int(math.Abs(float64(left.Height)-float64(right.Height))) <= 1)

		if left.Height > right.Height {
			return BalanceTree{Balance: balanced, Height: 1 + left.Height}
		}
		return BalanceTree{Balance: balanced, Height: 1 + right.Height}
	}

	return dfs(root).Balance
}
