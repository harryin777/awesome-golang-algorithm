package Solution

func Solution(x bool) bool {
	return x
}

func integerReplacement(n int) int {
	dp := make([]int, n+2)
	dp[0] = 1
	dp[1] = 0
	dp[2] = 1
	for i := 3; i < len(dp); i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2] + 1
		} else {
			dp[i] = dp[i-1] + 1
		}
	}

	for i := 3; i <= n; i++ {
		if i%2 != 0 {
			dp[i] = min(dp[i], dp[i+1]+1)
		}
	}

	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
