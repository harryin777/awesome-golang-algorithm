package Solution

/*
状态定义

	有股票 - 0
	无股票
		以前卖出、无冷冻 - 1
		今天卖出、有冷冻 - 2
		以前卖出、冷冻中 - 3

状态转移

	今天		昨天
	0
		<- 0
		<- 1
		<- 3
	1
		<- 1
		<- 3
	2
		<- 0
	3
		<- 2
*/
func maxProfit_1(prices []int) int {
	dp, n := make([][4]int, len(prices)), len(prices)
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3], dp[i-1][1])-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[n-1][3], max(dp[n-1][1], dp[n-1][2]))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// dp[i][k][n] i表示天数,k 表示交易次数,n 表示持有状态,0 是不持有, 1 是持有
func maxProfit_2(prices []int) int {
	dp := make([][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	//注意这里的 base case, 因为后面涉及到了 i-2 所以 i = 1 和 i = 0 的情况都要在这里列出来
	dp[0][1] = -prices[0]
	dp[0][0] = 0
	dp[1][1] = max(dp[0][1], -prices[0])
	dp[1][0] = max(dp[0][0], dp[0][1]-prices[0])
	for i := 2; i < len(dp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		// 注意这里, 冷冻期的意思是卖出股票以后隔一天才可以买入股票, 那么持有股票的状态要不是来自于前一天,要不是前两天的状态减去股票的价格
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i-1])
	}

	return dp[len(prices)][0]
}
