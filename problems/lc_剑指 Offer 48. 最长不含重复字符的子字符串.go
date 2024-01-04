package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s))
	dp[0] = 1
	lastIndexOf := make(map[uint8]int)
	lastIndexOf[s[0]] = 0
	for j := 1; j < len(s); j++ {
		// dp[j]定义为[0,j]的最长不重复子串
		// 状态转移方程:
		// 1. 当前遍历的字符如果没出现过,或者出现位置不在上一次的不重复子串中,说明可以在上次统计基础上增加一个长度
		if v, ok := lastIndexOf[s[j]]; !ok || v < j-dp[j-1] {
			dp[j] = dp[j-1] + 1
			// 2. 当前遍历的字符出现在了上一次的不重复子串中,则更新此次统计长度为当前下标j减去上次出现位置
		} else {
			dp[j] = j - lastIndexOf[s[j]]
		}

		lastIndexOf[s[j]] = j
	}

	return max(dp)
}

func max(nums []int) int {
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func main() {
	println(lengthOfLongestSubstring("au"))
}
