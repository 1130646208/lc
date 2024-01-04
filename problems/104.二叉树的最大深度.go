/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
var maxDepthVal *int

func maxDepth(root *TreeNode) int {
	maxDepthVal = new(int)
	*maxDepthVal = 0
	dfs(root, 0)
	return *maxDepthVal
}

func dfs(node *TreeNode, depth int) {
	if node == nil {
		if depth > *maxDepthVal {
			*maxDepthVal = depth
		}
		return
	}
	dfs(node.Left, depth+1)
	dfs(node.Right, depth+1)
}

// @lc code=end

