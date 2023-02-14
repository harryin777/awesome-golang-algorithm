package Solution

import "sort"

func Solution(amount int, coins []int) int {
	result := make([]int, amount+1)
	tmp := make([]int, amount+1)
	sort.Ints(coins)
	result[0] = 1
	baseCoin := coins[0]
	for idx := baseCoin; idx <= amount; idx += baseCoin {
		result[idx] = 1
	}
	for idx := 1; idx < len(coins); idx++ {
		for start := 0; start <= amount; start++ {
			for walker := 0; walker+start <= amount; walker += coins[idx] {
				tmp[walker+start] += result[start]
			}
		}
		for idx := 0; idx <= amount; idx++ {
			result[idx], tmp[idx] = tmp[idx], 0
		}
	}
	return result[amount]
}

// 这种解法重复了很多种情况,想想怎么剪枝
func change(amount int, coins []int) int {

	count := new(int)
	dp(amount, coins, count)
	return *count
}

func dp(amount int, coins []int, count *int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	for _, coin := range coins {
		res := dp(amount-coin, coins, count)
		if res == 0 {
			*count++
		} else {
			continue
		}
	}

	return *count
}

// dp[i][j] 前 i 个物品, amount 容量为 j 时, 最多的组合数
func change2(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				//第一种情况,不放进背包,那么和前 i-1 个的情况一样. 后一种,放进背包, 金额减去相应的值
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}
