package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	arrS, _ := scanner.ReadString('\n')
	arrS = strings.TrimSuffix(arrS, "\n")

	targetS, _ := scanner.ReadString('\n')
	targetS = strings.TrimSuffix(targetS, "\n")

	numsS := strings.Split(arrS, " ")

	var nums []int
	var target int
	var totalLess int

	target, _ = strconv.Atoi(targetS)

	for _, n := range numsS {
		vv, _ := strconv.Atoi(n)
		if vv < target {
			totalLess++
		}
		nums = append(nums, vv)
	}

	res := len(nums)

	// 滑动窗口,窗口大小是totalLess
	// 查看窗口里有多少大于target的,就是当前窗口所需的最小交换次数
	// 相当于穷举所有窗口
	for i := 0; i < len(nums)-totalLess+1; i++ {
		var tmpSwapTimes int
		for j := 0; j < totalLess; j++ {
			if nums[i+j] > target {
				tmpSwapTimes++
			}
		}
		if tmpSwapTimes < res {
			res = tmpSwapTimes
		}
	}

	fmt.Println(res)
}
