package tree

import (
	"fmt"
)

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	var dfs func(node *TreeNode, history string)

	dfs = func(n *TreeNode, h string) {
		if n == nil {
			return
		}
		s := fmt.Sprintf("%d", n.Val)

		if n.Left == nil && n.Right == nil {
			// backtrack
			res = append(res, h+s)
			return
		}

		dfs(n.Left, h+s+"->")
		dfs(n.Right, h+s+"->")
	}

	dfs(root, "")
	return res
}

func TestBinaryTreePaths() {
	s := []int{1, 2, 3, 5}
	root := CreateBST(s, 0, len(s)-1)
	fmt.Println(binaryTreePaths(root))
}
