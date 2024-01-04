/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(lNode *TreeNode, rNode *TreeNode) bool {
	if lNode == nil && rNode == nil {
		return true
	}
	if lNode == nil || rNode == nil {
		return false
	}
	return lNode.Val == rNode.Val &&
		compare(lNode.Right, rNode.Left) &&
		compare(lNode.Left, rNode.Right)
}

// @lc code=end

