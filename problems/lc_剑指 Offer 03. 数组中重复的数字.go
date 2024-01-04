package main

func findRepeatNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			j := nums[i]
			if nums[i] == nums[j] {
				return nums[i]
			}
			swap(nums, i, j)
		}
	}
	return 0
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
