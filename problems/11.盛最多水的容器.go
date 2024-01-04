/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	l, r := 0, len(height)-1

	var ret int
	for l < r {
		ret = max(min(height[l], height[r])*(r-l), ret)
		// 每次把较短的一个板子舍弃
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

