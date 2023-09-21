package tree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var dfs func(*TreeNode, int, map[int][]int)
	v := make(map[int][]int)

	maxDepth := 0
	dfs = func(n *TreeNode, lvl int, m map[int][]int) {
		if n == nil {
			return
		}
		if lvl > maxDepth {
			maxDepth = lvl
		}
		m[lvl] = append(m[lvl], n.Val)
		dfs(n.Left, lvl+1, m)
		dfs(n.Right, lvl+1, m)
	}

	dfs(root, 0, v)

	res := make([]int, maxDepth+1)

	for key, val := range v {
		res[key] = val[len(val)-1]
	}
	return res
}
