package Solution

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	// dp[i][j]代表 前nums中的前 i 个数字，能组成 j 的不同表达式的数目
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, target+1)
	}

	if nums[0] == 1 {
		dp[0][1] = 1
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= target; j++ {
			for k := i - 1; k < i; k++ {
				if k-1 < 0 {
					continue
				}
				dp[k][j] = dp[k-1][j] + dp[k-1][j-nums[k]] + 1
			}
		}
	}

	return dp[len(nums)][target]
}

func findTargetSumWays2(nums []int, target int) int {
	res := 0
	var dfs func(int, int)
	dfs = func(index int, count int) {
		if index == len(nums) {
			if count == target {
				res++
			}

			return
		}
		dfs(index+1, count+nums[index])
		dfs(index+1, count-nums[index])

	}
	dfs(0, 0)

	return res
}
