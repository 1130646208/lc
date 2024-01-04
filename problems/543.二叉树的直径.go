/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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

var ret *int

func diameterOfBinaryTree(root *TreeNode) int {
	ret = new(int)
	getDepth(root)
	return *ret
}

func getDepth(node *TreeNode) (int, int) {
	if node.Left == nil && node.Right == nil {
		return 0, 0
	}
	a, b := getDepth(node.Left)
	c, d := getDepth(node.Right)
	e, f := max(a+1, b+1), max(c+1, d+1)
	if e+f > *ret {
		*ret = e + f
	}
	return e, f
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

