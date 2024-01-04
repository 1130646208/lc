/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 维护这两个值就足够了
	var maxP int
	var minPrice int

	for i, p := range prices {
		if i == 0 {
			minPrice = p
		} else {
			minPrice = min(minPrice, p)
		}
		maxP = max(maxP, p-minPrice)
	}
	return maxP
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

