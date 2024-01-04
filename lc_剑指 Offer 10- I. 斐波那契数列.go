package main

func fib(n int) int {
	var a, b int64 = 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return int(a % 1000000007)
}

func main() {
	println(fib(45))
}
