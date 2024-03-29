package Solution

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	// dp 代表着i分钟内，从 0 到 j 点所需要的最小过路费
	dp := make([][]int, maxTime+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(passingFees))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = 1000 * 1001
		}

	}
	dp[0][0] = 0
	for i := 1; i <= maxTime; i++ {
		for _, edge := range edges {
			x, y, fee := edge[0], edge[1], edge[2]
			dp[i][y] = min(dp[i][y], dp[i-1][x]+fee)
		}
	}

	res := 1000 * 1001
	for _, val := range dp {
		res = min(res, val[len(edges)-1])
	}

	if res == 1000*1001 {
		return -1
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
