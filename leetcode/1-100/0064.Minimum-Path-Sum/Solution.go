package Solution

import "math"

// 这是自底向上
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			dp[i][0] = grid[i][0]
			for j := 1; j < n; j++ {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// 这是自顶向下的暴力递归
func dp(grid [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	if i < 0 || j < 0 {
		return int(math.MaxInt32)
	}
	return Min(
		dp(grid, i-1, j),
		dp(grid, i, j-1),
	) + grid[i][j]
}
