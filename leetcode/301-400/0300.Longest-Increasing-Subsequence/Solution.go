package Solution

// dp[i]的定义是以 i 为结尾的最长递增子序列
func lengthOfLIS1(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	ans := -1
	for i := 1; i < len(dp); i++ {
		for j := 1; j < i; j++ {
			if nums[i-1] > nums[j-1] {
				dp[i] = max(dp[i], dp[j]+1)
				ans = max(ans, dp[i])
			}
		}
	}

	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
