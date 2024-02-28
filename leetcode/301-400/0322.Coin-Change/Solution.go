package Solution

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// 动态规划（不压缩空间）
func coinChange2(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1) // 定义状态
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}
	// 初始条件
	for j := 0; j <= amount; j++ {
		dp[0][j] = amount + 1
	}
	dp[0][0] = 0
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			// 状态转移方程
			if j >= coins[i-1] {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-coins[i-1]]+1)))
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	if dp[len(coins)][amount] > amount {
		return -1
	} else {
		return dp[len(coins)][amount]
	}
}

// 动态规划（压缩空间）
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始值
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-coins[i-1]]+1)))
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

func CoinChange4(coins []int, amount int) int {

	dpMap := make(map[int]int)

	return dpCoin(coins, amount, dpMap)
}

var mm = int(^uint(0) >> 1)

func dpCoin(coins []int, amount int, dpMap map[int]int) int {
	if val, exists := dpMap[amount]; exists {
		return val
	}

	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	val := mm
	for i := 0; i < len(coins); i++ {
		res := dpCoin(coins, amount-coins[i], dpMap)
		if res == -1 {
			continue
		}
		val = min(val, res+1)
		dpMap[amount] = val
	}
	if val == mm {
		return -1
	}

	return val
}

// 为啥 dp 数组中的值都初始化为 amount + 1 呢，因为凑成 amount 金额的硬币数最多只可能等于 amount（全用 1 元面值的硬币）.
func CoinChange5(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coins[j]]+1)))
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

// dp[i] 凑成i 金额的硬币,的最少个数
func coinChange4(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	val := mm
	for _, coin := range coins {
		res := coinChange4(coins, amount-coin)
		if res < 0 {
			continue
		}
		val = min(val, res+1)
	}
	if val == mm {
		return -1
	}

	return val
}

func coinChange6(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = j
			}
		}
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-coins[i-1]]+1)
			}
		}
	}

	if dp[len(coins)][amount] == 0 {
		return -1
	}

	return dp[len(coins)][amount]
}

// 这里自底向上的方法, 是从最基础的情况一步步向 amount 逼近的,所以虽然最终没有把当前数字,i 减为 0,实际上在 i 的子问题里已经构成了 i
func coinChange7(coins []int, amount int) int {
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, c := range coins {
			if i >= c && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func coinChange8(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	res := math.MaxInt32
	for _, coin := range coins {
		subProblem := coinChange8(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		res = min(res, subProblem+1)
	}

	if res == math.MaxInt32 {
		return -1
	}

	return res
}

func coinChange9(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make(map[int]int)
	var dfs func(coins []int, balance int) int
	dfs = func(coins []int, balance int) int {
		if balance < 0 {
			return -1
		}
		if balance == 0 {
			return 0
		}
		if val, ok := dp[amount]; ok {
			return val
		}
		res := int(^uint(0) >> 1)
		for i := 0; i < len(coins); i++ {
			ans := dfs(coins, balance-coins[i])
			if ans < 0 {
				continue
			}
			res = min(res, ans+1)
		}
		dp[amount] = res
		return res
	}

	return dfs(coins, amount)
}
