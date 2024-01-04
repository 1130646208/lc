package main

import "fmt"

func decodeString(s string) string {
	var factor int
	var res []rune
	var resStack [][]rune
	var factorStack []int
	for _, r := range s {
		if ok, val := isDigit(r); ok {
			factor += factor*10 + val
		} else if isLetter(r) {
			res = append(res, r)
		} else if r == '[' {
			resStack = append(resStack, res)
			factorStack = append(factorStack, factor)
			factor = 0
			res = []rune{}
		} else if r == ']' {
			resPeek := resStack[len(resStack)-1]
			resStack = resStack[:len(resStack)-1]
			factorPeek := factorStack[len(factorStack)-1]
			factorStack = factorStack[:len(factorStack)-1]
			tmp := make([]rune, len(res))
			copy(tmp, res)
			for i := 0; i < factorPeek; i++ {
				res = append(res, tmp...)
			}
			res = append(res, resPeek...)
		} else {
			panic("invalid")
		}
	}
	return string(res)
}

func isDigit(r rune) (bool, int) {
	res := r - '0'
	return res >= 0 && res <= 9, int(res)
}

func isLetter(r rune) bool {
	return r-'a' >= 0 && r-'a' <= 25
}
func main() {
	s := decodeString("3[a]2[bc]")
	fmt.Print(s)
}
