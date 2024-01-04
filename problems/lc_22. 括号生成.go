package main

func generateParenthesis(n int) []string {
	var path []rune
	var ret []string
	comb(0, 0, path, &ret, n)
	return ret
}

func comb(lNum, rNum int, path []rune, ret *[]string, combNum int) {
	if lNum > combNum || rNum > combNum {
		return
	}
	if lNum == rNum && lNum == combNum {
		tmp := make([]rune, lNum)
		copy(tmp, path)
		*ret = append(*ret, string(path))
	}

	// 这里不用for循环
	path = append(path, '(')
	comb(lNum+1, rNum, path, ret, combNum)
	path = path[:len(path)-1]
	if lNum > rNum {
		path = append(path, ')')
		comb(lNum, rNum+1, path, ret, combNum)
		path = path[:len(path)-1]
	}
}
