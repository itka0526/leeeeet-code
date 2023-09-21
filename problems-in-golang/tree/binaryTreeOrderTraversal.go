package tree

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var level int

	temp := make([]*TreeNode, 1)
	temp[0] = root

	for len(temp) > 0 {
		currLength := len(temp)
		for i := 0; i < currLength; i++ {
			node := temp[0]
			temp = temp[1:]

			if len(result) <= level {
				result = append(result, []int{node.Val})
			} else {
				result[level] = append(result[level], node.Val)
			}

			if node.Left != nil {
				temp = append(temp, node.Left)
			}

			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		level++
	}

	return result
}
