package Solution

/*
操作

	dp[i][0] 没有操作

	dp[i][1] 第一次买入
		以前已经买入｜以前未买入、今天买入
	dp[i][2] 第一次卖出
		以前已经卖出｜以前未卖出、今天卖出
	dp[i][3] 第二次买入
		...
	dp[i][4] 第二次卖出
		...
*/
func maxProfit_1(prices []int) int {
	dp, n := make([][5]int, len(prices)), len(prices)
	dp[0][1], dp[0][3] = -prices[0], -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[n-1][4]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// dp[i][k][n] i是天数,k 是当前状态还剩余的交易次数,n 是是否持有股票,0 不持有,1 持有
// k理解成 fee 的概念,都是在买卖操作的时候扣减 1
func maxProfit_2(prices []int) int {
	dp := make([][][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j <= 2; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][0] = -prices[i]
				continue
			}
			if j == 0 {
				dp[i][j][0] = dp[i-1][j][0]
				dp[i][j][1] = dp[i-1][j][1]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}

	return dp[len(prices)][2][0]
}
