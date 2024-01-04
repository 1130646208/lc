/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	var needMap = make(map[uint8]int)
	var needCount int
	for _, ch := range t {
		needMap[uint8(ch)]++
		needCount++
	}

	var ret string
	var i int
	for j := i; j < len(s); j++ {
		for k := i; k <= j; k++ {
			if needMap[s[i]] > 0 {
				needMap[s[i]]--
				needCount--
				if needCount == 0 && (len(ret) == 0 || j-i > len(ret)) {
					ret = s[i : j+1]
				}
			}
		}
	}

	fmt.Println(ret)
	return ret
}

// @lc code=end

