package main

import (
	"fmt"
)

func main() {
	var i int
	for {
		num, _ := fmt.Scanf("%d", &i)
		if num == 0 {
			break
		}
	}
	println(numOf1(i))
}

func numOf1(n int) int {
	var numOf1s, maxContinues int

	for i := 0; ; i++ {
		if 1<<i > n {
			break
		}
		res := n & (1 << i)
		if res == 1<<i {
			numOf1s++
			if maxContinues < numOf1s {
				maxContinues = numOf1s
			}
		} else {
			numOf1s = 0
		}
	}

	return maxContinues
}
