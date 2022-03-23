package main

//10-1 斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
	// return fib(n) + fib(n-1)
}

//10-II 青蛙跳台阶

func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	a, b, t := 1, 1, 0
	for i := 2; i <= n; i++ {
		t = a + b
		a = b
		b = t
	}
	const mod int = 1e9 + 7
	return t % mod
}
