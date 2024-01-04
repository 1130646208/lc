package main

import (
	"math"
)

func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	var minus bool
	// same
	if a > 0 && b > 0 || a < 0 && b < 0 {
		minus = false
	} else {
		minus = true
	}

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	var res int
	for a >= b {
		var i int
		for factor := 1; ; factor <<= i {
			a -= b * factor
			if a < 0 {
				a += b * factor
				break
			} else if a == 0 {
				res += factor
				break
			}
			res += factor
			i++
		}
	}
	if minus {
		return -res
	}
	return res
}
func main() {
	println(divide(-90, 69))
}
