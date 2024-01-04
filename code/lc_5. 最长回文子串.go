package main

import "fmt"

func longestPalindrome(s string) string {
	var ret string
	ret = string(s[0])
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	for i := 0; i < len(s)-1; i++ {
		for j := i; j > i && j < len(s); j++ {
			if s[i] == s[j] {
				if j-i == 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] && j-i+1 > len(ret) {
				ret = s[i : j+1]
			}
		}
	}
	return ret
}

func main() {
	s := longestPalindrome("abcba")
	fmt.Print(s)
}
