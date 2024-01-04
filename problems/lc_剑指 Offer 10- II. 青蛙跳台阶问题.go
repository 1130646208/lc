package main

var a = make(map[int]int)

func numWays(n int) int {
	var res int
	if n == 1 || n == 0 {
		return 1
	}
	if n == 2 {
		return 2
	}

	if _, ok := a[n]; ok {
		return a[n]
	} else {
		a[n] = (numWays(n-1) + numWays(n-2)) % 1000000007
		res = a[n]
	}

	return res
}

func main() {
	println(numWays(3))
}
