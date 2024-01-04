package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	index := -1

	for l <= r {
		mid := (r + l) >> 1
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			index = mid
			break
		}
	}

	if index == -1 {
		return 0
	}

	sum := 0
	for i := index; i > -1 && nums[i] == target; i-- {
		sum++
	}

	for i := index; i < len(nums) && nums[i] == target; i++ {
		sum++
	}

	return sum - 1
}

func main() {
	search([]int{5, 7, 7, 8, 8, 10, 11}, 5)
}
