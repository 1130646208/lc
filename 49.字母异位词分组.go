/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	var m = make(map[[26]int][]string)
	var ret = [][]string{}
	for _, word := range strs {
		key := [26]int{}
		for _, r := range word {
			key[r-'a'] += 1
		}
		m[key] = append(m[key], word)
	}

	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

// @lc code=end

