package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return comp(root.Left, root.Right)
}

func comp(A, B *TreeNode) bool {
	if A != nil && B == nil {
		return false
	} else if A == nil && B != nil {
		return false
	} else if A == nil && B == nil {
		return true
	}
	return A.Val == B.Val && comp(A.Left, B.Right) && comp(A.Right, B.Left)
}
