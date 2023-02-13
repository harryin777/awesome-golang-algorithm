package Solution

// dp[i][j] 表示对于前 i 个数字,可以刚好组成j
func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}

	if sum%2 != 0 {
		return false
	}

	volume := sum / 2
	dp := make([][]bool, len(nums)+1)
	for i := 0; i <= len(dp); i++ {
		dp[i] = make([]bool, volume+1)
		dp[i][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		dp[i] = make([]bool, volume+1)
		for j := 1; j <= volume; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[len(nums)][sum]
}
