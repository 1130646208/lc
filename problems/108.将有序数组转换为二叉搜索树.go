/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, l, r int) *TreeNode {
	if r < l {
		return nil
	}
	mid := (l + r) / 2
	root := new(TreeNode)
	root.Val = nums[mid]
	root.Left = dfs(nums, l, mid-1)
	root.Right = dfs(nums, mid+1, r)
	return root
}

// @lc code=end

