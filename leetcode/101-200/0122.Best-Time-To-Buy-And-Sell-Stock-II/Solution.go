package Solution

func maxProfit_1(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

/*
当前有无股票

	有
	    昨天有、今天不动
	    昨天无、今天买
	无
	    昨天有、今天卖
	    昨天无、今天不动
*/
func maxProfit_2(prices []int) int {
	dp, n := make([][2]int, len(prices)), len(prices)
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[n-1][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// dp[i][k][n] i是天数,k 是当前剩余交易次数,n 是当前是否持有股票,0 不持有,1 持有, k为正无穷,还是可以忽略掉
func maxProfit_3(prices []int) int {
	dp := make([][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	//注意这里的 base case
	dp[0][1] = -prices[0]
	dp[0][0] = 0

	for i := 1; i < len(dp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		// 因为这里k 是正无穷,也就是可以有之前的交易,所以才有i-1
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1])
	}

	return dp[len(prices)][0]
}
