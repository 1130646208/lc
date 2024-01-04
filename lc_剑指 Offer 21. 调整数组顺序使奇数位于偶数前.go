package main

func exchange(nums []int) []int {

	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]&1 == 1 {
			i++
		}
		if nums[j]&1 == 0 {
			j--
		}
		if nums[j]&1 == 0 && nums[i]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
