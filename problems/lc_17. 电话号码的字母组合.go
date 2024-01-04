package main

func LetterCombinations(digits string) []string {
	var ret []string

	var keyMapping = map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var path []rune
	var keys [][]rune
	for _, r := range digits {
		keys = append(keys, keyMapping[r])
	}
	permuate(path, keys, 0, &ret)
	return ret
}

func permuate(path []rune, keys [][]rune, ind int, ret *[]string) {
	if ind == len(keys) {
		tmp := make([]rune, len(path))
		copy(tmp, path)
		*ret = append(*ret, string(path))
		return
	}

	for _, r := range keys[ind] {
		path = append(path, r)
		permuate(path, keys, ind+1, ret)
		path = path[:len(path)-1]
	}
}
