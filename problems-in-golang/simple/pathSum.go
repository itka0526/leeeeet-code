package simple

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(curr *TreeNode, currSum int, targetSum int) bool {
	if curr == nil {
		return false
	} else if curr.Left == nil && curr.Right == nil {
		return (currSum + curr.Val) == targetSum
	} else {
		return dfs(curr.Left, currSum+curr.Val, targetSum) || dfs(curr.Right, currSum+curr.Val, targetSum)
	}
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs(root, 0, targetSum)
}
