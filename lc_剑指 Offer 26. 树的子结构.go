package main

// 此函数拿着B树往A树上的结点一个一个试
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 如果B是个空树，根据题干要求，空树不是任何数的子结构
	if B == nil {
		return false
		// 如果比到了越过A的叶子结点，说明都没找到匹配的
	} else if A == nil && B != nil {
		return false
	}

	// 如果当前比较两个根节点没匹配，继续拿着B往A树的左右子树一个一个试
	if A.Val != B.Val {
		return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
	}
	// 如果当前比较两个根节点匹配了，不仅要递归比较两个数的左右子树是否相同，继续拿着B往A树的左右子树一个一个试
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(A *TreeNode, B *TreeNode) bool {
	if A != nil && B == nil {
		return true
	} else if B != nil && A == nil {
		return false
	} else if A == nil && B == nil {
		return true
	}
	return A.Val == B.Val && recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
