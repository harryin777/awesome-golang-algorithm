package Solution

func maxProfit_1(prices []int) int {
	ans := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			ans = max(ans, prices[j]-prices[i])
		}
	}
	return ans
}

func maxProfit_2(prices []int) int {
	ans, lowPrice := 0, prices[0]
	for i := 0; i < len(prices); i++ {
		ans = max(ans, prices[i]-lowPrice)
		lowPrice = min(lowPrice, prices[i])
	}
	return ans
}

/*
当前有无股票

	有
	    昨天有、今天不动
	    昨天无、今天买【第一次买】
	无
	    昨天有、今天卖
	    昨天无、今天不动
*/
func maxProfit_3(prices []int) int {
	dp, n := make([][2]int, len(prices)), len(prices)
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[n-1][1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var mm = ^int(^uint(0) >> 1)

// dp[i][k][n] i是天数,k 是当前剩余交易次数,n 是交易状态, 本题 k 恒为 1 忽略掉这个维度
func maxProfit4(prices []int) int {
	dp := make([][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2) // 两种状态
	}

	dp[0][1] = -prices[0]
	dp[0][0] = 0

	for i := 1; i < len(dp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		// 注意这里为什么是直接减去了 price,因为只能进行一次操作,所以没有 i-1 的状态了,如果有,那么当前的交易是不会发生的
		dp[i][1] = max(dp[i-1][1], -prices[i-1])
	}

	if dp[len(prices)][0] < 0 {
		return 0
	}

	return dp[len(prices)][0]
}
