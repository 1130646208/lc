package main

func translateNum(num int) int {
	nums := []int{}
	for num != 0 {
		n := num % 10
		nums = append(nums, n)
		num /= 10
	}
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 1, 1

	// 类似于青蛙爬楼梯
	// 当当前位和上一位不能组合翻译时,则说明不算这一位的时候能翻译多少种,算上这一位还是没增加翻译数量
	// 当当前这一位和上一位可以组合翻译时,则说明总翻译数量=除去当前位翻译数+除去当前位和上一位翻译数
	for i := len(nums) - 2; i >= 0; i-- {
		if s := nums[i] + nums[i+1]*10; s < 26 && s > 9 {
			dp[len(nums)-i] = dp[len(nums)-i-1] + dp[len(nums)-i-2]
		} else {
			dp[len(nums)-i] = dp[len(nums)-i-1]
		}
	}
	return dp[len(nums)-1]
}

func main() {
	println(translateNum(12258))
}
