package main

func maxSubArray(nums []int) int {

	dp := make([]int, len(nums)+1)
	var max, ind int = dp[0], 0

	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + nums[i-1]
		if dp[i] > max {
			max = dp[i]
			ind = i
		}
	}

	var min int = dp[ind]
	for i := ind; i >= 0; i-- {
		if dp[i] < min {
			min = dp[i]
		}
	}

	var max2 int = nums[0]
	// all minus
	if ind == 0 {
		for i := 0; i < len(nums); i++ {
			if nums[i] > max2 {
				max2 = nums[i]
			}
		}
		return max2
	}

	return max - min
}

func main() {
	print(maxSubArray([]int{-2, -1}))
}
