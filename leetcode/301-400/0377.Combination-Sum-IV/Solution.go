package Solution

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func combinationSum42(nums []int, target int) int {
	var dfs func(curr int)
	res := 0
	dfs = func(curr int) {
		if curr > target {
			return
		}
		if curr == target {
			res++
			return
		}
		for i := 0; i < len(nums); i++ {
			curr += nums[i]
			dfs(curr)
			curr -= nums[i]
		}
	}
	dfs(0)
	return res
}
