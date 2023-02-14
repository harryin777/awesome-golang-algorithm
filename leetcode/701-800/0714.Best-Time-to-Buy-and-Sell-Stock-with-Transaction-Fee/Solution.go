package Solution

func maxProfit_1(prices []int, fee int) int {
	ans, minPrice := 0, prices[0]
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v > minPrice && v > minPrice+fee {
			ans += v - minPrice - fee
			minPrice = v - fee
		}
	}
	return ans
}

func maxProfit_2(prices []int, fee int) int {
	dp, n := make([][2]int, len(prices)), len(prices)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit_3(prices []int, fee int) int {
	dp := make([][]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(dp); i++ {
		// 注意这里的费用可以在买入的时候减去,也可以在卖出的时候减去,但是只能减去一次
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1])
	}

	return dp[len(prices)][0]
}
