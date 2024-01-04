package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	ind := bisearch(nums, target, 0, len(nums)-1)
	l, r := -1, -1
	if ind == -1 {
		return []int{l, r}
	}
	for x := ind; x < len(nums); x++ {
		if nums[x] == target {
			r = x
		}
	}
	for x := ind; x > -1; x-- {
		if nums[x] == target {
			l = x
		}
	}
	return []int{l, r}
}

func bisearch(nums []int, target int, l, r int) int {
	if len(nums) == 0 || r < l {
		return -1
	}
	mid := (l + r) / 2
	if target > nums[mid] {
		return bisearch(nums, target, mid+1, r)
	} else if target < nums[mid] {
		return bisearch(nums, target, l, mid-1)
	} else {
		return mid
	}
}

func main() {
	nums := []int{2, 2}
	fmt.Println(bisearch(nums, 3, 0, len(nums)-1))
}
