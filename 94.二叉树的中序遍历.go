/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var ret []int

	stack := &[]*TreeNode{}

	// 注意 || 是关键
	for len(*stack) > 0 || root != nil {
		for root != nil {
			push(root, stack)
			root = root.Left
		}
		root = pop(stack)
		// 判空是关键
		if root != nil {
			ret = append(ret, root.Val)
			root = root.Right
		}
	}

	return ret
}

func push(ele *TreeNode, stack *[]*TreeNode) {
	*stack = append(*stack, ele)
}

func pop(stack *[]*TreeNode) *TreeNode {
	l := len(*stack) - 1
	ret := (*stack)[l]
	*stack = (*stack)[0:l]
	return ret
}

// @lc code=end

