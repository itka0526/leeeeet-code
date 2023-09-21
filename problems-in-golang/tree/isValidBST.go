package tree

import "math"

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt, math.MaxInt)
}

func isValid(node *TreeNode, left, right int) bool {
	if node == nil {
		return true
	}

	if !(left < node.Val && node.Val < right) {
		return false
	}

	return isValid(node.Left, left, node.Val) && isValid(node.Right, node.Val, right)
}
