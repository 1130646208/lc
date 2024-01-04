/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	// 存储某个值所处连续区域最长连续数字个数(只有边界元素才有意义)
	var numMap = make(map[int]int)
	var mContinuous int

	for _, n := range nums {
		// 只处理map中没有的元素,没有的一定是边界元素
		if _, ok := numMap[n]; !ok {
			// 边界元素的左右两边(如果已经加入了map)一定也都是边界元素
			l := numMap[n-1]
			r := numMap[n+1]

			curLen := l + r + 1
			if curLen > mContinuous {
				mContinuous = curLen
			}
			// 更新本元素以及其两边的边界元素
			numMap[n] = curLen
			numMap[n-l] = curLen
			numMap[n+r] = curLen
		}
	}
	return mContinuous
}

// @lc code=end

