package Solution

func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}

	step1, step2 := 0, 1
	ans := 0
	for i := 2; i <= n; i++ {
		ans = step1 + step2
		step1 = step2
		step2 = ans
	}

	return ans
}
