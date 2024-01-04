package main

func countBits(n int) []int {
	// dp[i] = dp[i-1] + 1; i 是奇数
	// dp[i] = dp[i/2];     i 是偶数

	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i&1 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i-1] + 1
		}
	}
	return res
}
