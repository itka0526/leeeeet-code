package tree

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	left := maxDepth(root.Left)
// 	right := maxDepth(root.Right)
// 	if left > right {
// 		return left + 1
// 	}
// 	return right + 1
// }

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	left := maxDepth(root.Left)
// 	right := maxDepth(root.Right)
// 	if left > right {
// 		return left + 1
// 	}
// 	return right + 1
// }

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth1(root.Left)
	right := maxDepth1(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
