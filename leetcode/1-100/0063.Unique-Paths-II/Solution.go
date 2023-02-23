package Solution

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m, n int

	if m = len(obstacleGrid); m == 0 {
		return 0
	}
	if n = len(obstacleGrid[0]); n == 0 {
		return 0
	}

	dp := [][]int{}
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}
	dp[m][n-1] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([][]int, m)
	column := false
	_ = column
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		if column {
			dp[i][0] = 0
			continue
		}
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			column = true
			dp[i][0] = 0
		}
	}

	row := false
	_ = row
	for i := 0; i < len(dp[0]); i++ {
		if row {
			dp[0][i] = 0
			continue
		}
		if obstacleGrid[0][i] == 0 {
			dp[0][i] = 1
		} else {
			row = true
			dp[0][i] = 0
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// if obstacleGrid[i-1][j] == 0 && obstacleGrid[i][j-1] == 0{
			//     dp[i][j] = dp[i-1][j] + dp[i][j-1]
			// }
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
