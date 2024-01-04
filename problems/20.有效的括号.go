/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {

	stack := []rune{}

	for _, ch := range s {
		switch ch {
		case '[', '(', '{':
			stack = append(stack, ch)
		default:
			// 记得判断这里，否则如果第一个是右括号会报错
			if len(stack)-1 < 0 {
				return false
			}
			tmp := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if !match(tmp, ch) {
				return false
			}
		}
	}
	return len(stack) == 0
}

func match(l, r rune) bool {
	m := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	return m[l] == r
}

// @lc code=end

