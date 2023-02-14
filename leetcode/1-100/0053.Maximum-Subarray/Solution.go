package Solution

func maxSubArray(nums []int) int {
	dp, ans := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		ans = max(dp[i], ans)
		// fmt.Println(i, ans, dp, nums)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray2(nums []int) int {
	return calc(0, len(nums)-1, nums)
}

func calc(l, r int, nums []int) int {
	if l == r {
		return nums[l]
	}
	mid := (l + r) >> 1
	lmax, lsum, rmax, rsum := nums[mid], 0, nums[mid+1], 0

	for i := mid; i >= 1; i-- {
		lsum += nums[i]
		lmax = max(lmax, lsum)
	}

	for i := mid + 1; i <= r; i++ {
		rsum += nums[i]
		rmax = max(rmax, rsum)
	}

	return max(max(calc(l, mid, nums), calc(mid+1, r, nums)), lmax+rmax)
}

func maxSubArray3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = nums[i]
	}

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i], dp[i-1]+dp[i])
	}

	res := ^int(^uint(0) >> 1)
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}

	return res
}

func maxSubArray4(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp0 := nums[0]
	dp1, res := 0, dp0

	for i := 1; i < len(nums); i++ {
		dp1 = max(nums[i], dp0+nums[i])
		dp0 = dp1
		res = max(res, dp1)
	}

	return res
}

var minOne = ^int(^uint(0) >> 1)

// dp[i]以 num[i] 结尾的最大连续子数组的和
func maxSubArray31(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = nums[i]
	}

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i], dp[i-1]+nums[i])
	}

	res := minOne
	for _, val := range dp {
		res = max(res, val)
	}

	return res
}
