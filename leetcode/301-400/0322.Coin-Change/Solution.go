package Solution

import (
	"math"
	"sort"
)

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	sort.Ints(coins)
	for i := 1; i < len(dp); i++ {
		for j := len(coins) - 1; j >= 0; j-- {
			if i < coins[j] {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
