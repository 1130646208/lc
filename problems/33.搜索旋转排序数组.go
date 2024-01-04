/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {

	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + ((r - l) >> 1)
		if nums[mid] == target {
			return mid
		}
		// 先判断mid在断点左边还是右边
		if nums[mid] >= nums[l] { // 重要
			// mid在左边
			// 再判断target在mid左边还是右边
			if target >= nums[l] && target < nums[mid] {
				// target在mid左边
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// mid在右边
			// 再判断target在mid左边还是右边
			if target > nums[mid] && target <= nums[r] {
				// target在mid右边
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end

