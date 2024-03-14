package Solution

func Solution(x bool) bool {
	return x
}

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]uint8, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]uint8, n+1)
		if i > 0 {
			dp[i][0] = s1[i-1]
		}
	}

	for i := 0; i < len(dp[0]); i++ {
		if i > 0 {
			dp[0][i] = s2[i-1]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s1[i] == s2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j]) + min(s1[i], s2[j])
			}
		}
	}

	return int(dp[m][n])
}

func min(x, y uint8) uint8 {
	if x > y {
		return y
	}
	return x
}
