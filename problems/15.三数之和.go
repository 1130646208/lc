/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	// 此题去重是难点
	// 容易重复的例子[-2,0,0,2,2]
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ret := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		if nums[i] > 0 {
			continue
		}
		// i指针去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < k {
			res := nums[i] + nums[j] + nums[k]
			if res > 0 && k-1 > j {
				k--
			} else if res < 0 && j+1 < k {
				j++
			} else if res == 0 {
				// k指针去重
				if k < len(nums)-1 && nums[k] == nums[k+1] {
					k--
					continue
				}
				// j指针去重
				if j > i+1 && nums[j] == nums[j-1] {
					j++
					continue
				}
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else {
				break
			}
		}
	}
	return ret
}

// @lc code=end

