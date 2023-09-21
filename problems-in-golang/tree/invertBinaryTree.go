// func invertTree1(root *TreeNode) *TreeNode {
// 	if root != nil {
// 		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
// 	}
// 	return root
// }

// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	left := invertTree(root.Left)
// 	right := invertTree(root.Right)

// 	root.Left = right
// 	root.Right = left
// 	return root
// }

// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	left := invertTree(root.Left)
// 	right := invertTree(root.Right)

// 	root.Left = right
// 	root.Right = left

// 	return root
// }

// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	left := invertTree(root.Left)
// 	right := invertTree(root.Right)

// 	root.Left = right
// 	root.Right = left

//		return root
//	}

package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func invertTree(root *TreeNode) *TreeNode {
// 	if root != nil {
// 		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
// 	}
// 	return root
// }

// func invertTree(root *TreeNode) *TreeNode {
// 	if root != nil {
// 		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
// 	}

// 	return root
// }

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}
