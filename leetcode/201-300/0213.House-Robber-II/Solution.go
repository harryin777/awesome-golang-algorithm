package Solution

func Solution(x bool) bool {
	return x
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	l := len(nums)
	dp := make([]int, l-1)
	dp2 := make([]int, l-1)
	dp[0] = nums[0]
	dp[1] = max(nums[1], dp[0])

	//这里可以优化,只使用了 i, i-1, i-2 所以你懂的
	for i := 2; i < l-1; i++ {
		dp[i] = max(dp[i], max(dp[i-1], dp[i-2]+nums[i]))
	}

	nums = nums[1:]
	dp2[0] = nums[0]
	dp2[1] = max(nums[1], dp2[0])

	for i := 2; i < l-1; i++ {
		dp2[i] = max(dp2[i], max(dp2[i-1], dp2[i-2]+nums[i]))
	}

	return max(dp[l-2], dp2[l-2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
