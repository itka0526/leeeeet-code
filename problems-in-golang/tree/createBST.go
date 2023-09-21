package tree

func CreateBST(s []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	node := &TreeNode{Val: s[mid], Left: CreateBST(s, start, mid-1), Right: CreateBST(s, mid+1, end)}
	return node
}
