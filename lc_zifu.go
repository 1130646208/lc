package main

import (
	"fmt"
	"strings"
)

func strToInt(str string) int {
	str = strings.TrimLeft(str, " ")
	var minus int32
	var res int32
	var nums []int32
	for i, ch := range str {
		if i == 0 {
			if ch == '+' {
				minus = 1
				continue
			} else if ch == '-' {
				minus = -1
				continue
			} else {
				minus = 1
			}
		}
		if n := ch - '0'; n >= 0 && n < 10 {
			nums = append(nums, n)
		} else {
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if tmpRes := (res*10 + nums[i]) / 10; tmpRes == res {
			res = res*10 + nums[i]
		} else if minus == -1 {
			return -(1 << 31)
		} else if minus == 1 {
			return 1<<31 - 1
		}
	}
	return int(res * minus)
}

func main() {
	i := strToInt("-u6147483648")
	fmt.Println(i)
}
