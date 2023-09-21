package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// running through ideas
	// if p is found or q is found
	// try to go backtrace and return common ancestor

	// the solution is to trace through the tree
	// got the basic idea

	// the root will always be the common ancestor and we just need to go deeper into the tree to find least common ancestor

	// explaining it in code seems apropriate

	if p.Val > root.Val && q.Val > root.Val {
		// we must dive deeper to the right of the tree because p and q have common ancestor on that side
		return lowestCommonAncestor(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		// same goes here too !
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}
