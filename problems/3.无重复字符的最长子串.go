/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// ret[i]记录以s[i]结尾的最长不重复子串的长度
	var ret = make([]int, len(s))
	// mem记录某个字符最后出现的下标
	var mem = make(map[rune]int)
	var result int
	for i, ch := range s {
		if _, ok := mem[ch]; !ok {
			// 没出现过,则说明会贡献一个长度
			if i == 0 {
				ret[i] = 1
			} else {
				ret[i] = ret[i-1] + 1
			}
		} else {
			// 若出现过,则看下上次是在哪里(mem[ch])出现的
			// i-mem[ch]代表当前位置到上次出现的位置有多远
			// 假设当前字符没出现过,则往回倒ret[i-1]+1个字符,看看是否会包含到当前字符
			// 如果包含到当前字符了,则肯定不能+1了,只能算到当前字符上次出现的位置的后一位
			ret[i] = min(ret[i-1]+1, i-mem[ch])
		}
		// 每次都记录最后一个下标
		mem[ch] = i
		// 更新结果
		if ret[i] > result {
			result = ret[i]
		}
	}
	return result
}

// @lc code=end

