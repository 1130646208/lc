/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
var rem = make(map[int]int)

func climbStairs(n int) int {
	if ret, ok := rem[n]; ok {
		return ret
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	res := climbStairs(n-1) + climbStairs(n-2)
	rem[n] = res
	return res
}

// @lc code=end

