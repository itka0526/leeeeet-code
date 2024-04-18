package medium

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, path string, paths *[]string) {
	if node.Left == nil && node.Right == nil {
		t := string(rune(node.Val) + 'a')
		for i := len(path) - 1; i >= 0; i-- {
			t += string(path[i])
		}
		*paths = append(*paths, t)
		return
	}
	path += string(rune(node.Val) + 'a')
	if node.Left != nil {
		dfs(node.Left, path, paths)
	}
	if node.Right != nil {
		dfs(node.Right, path, paths)
	}
}

func smallestFromLeaf(root *TreeNode) string {
	ans := []string{}
	dfs(root, "", &ans)
	sort.Strings(ans)
	fmt.Println(ans)
	return ans[0]
}
