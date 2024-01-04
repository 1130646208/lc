/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	// 摩尔投票法
	var vote = 0
	var ret int
	for _, n := range nums {
		if vote == 0 {
			ret = n
			vote++
			continue
		}
		if n == ret {
			vote++
		} else {
			vote--
		}
	}
	return ret
}

// @lc code=end

