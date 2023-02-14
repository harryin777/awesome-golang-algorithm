package Solution

// 普通的动态规划
func rob(nums []int) int {
	dp := make([][2]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = nums[i-1] + dp[i-1][0]
	}
	// fmt.Println(dp)
	return max(dp[len(nums)][0], dp[len(nums)][1])
}

// 优化空间
func rob2(nums []int) int {
	prevNo, prevYes := 0, 0

	for _, v := range nums {
		tmp := prevNo
		prevNo = max(prevNo, prevYes)
		prevYes = v + tmp
	}
	return max(prevNo, prevYes)
}

func rob3(nums []int) int {
	robEven, robOdd := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			robEven = max(robEven+nums[i], robOdd)
		} else {
			robOdd = max(robEven, nums[i]+robOdd)
		}
	}
	return max(robEven, robOdd)
}

func rob4(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// dp[i] 前 i 个数字 两数字的最大和
func rob11(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[1], dp[0])

	//这里可以优化,只使用了 i, i-1, i-2 所以你懂的
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i], max(dp[i-1], dp[i-2]+nums[i]))
	}

	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}

	return res
}

// 这种才是 dp[i] i 是以 num[i] 结尾的做法
func rob12(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]

	for i := 2; i < len(nums); i++ {
		val := 0
		for j := 0; j < i-1; j++ {
			val = max(val, dp[j])
		}
		dp[i] = val + nums[i]
	}

	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}

	return res
}
