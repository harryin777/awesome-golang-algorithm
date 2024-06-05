package Solution

func wiggleMaxLength(nums []int) int {
	//	序列个数小于二直接是摇摆序列
	if len(nums) < 2 {
		return len(nums)
	}
	//	3种状态
	const BEGIN, UP, DOWN = 0, 1, 2
	//	摇摆序列最小长度至少为1
	STATE, maxLength := BEGIN, 1

	for i := 1; i < len(nums); i++ {
		switch STATE {
		case BEGIN:
			if nums[i-1] < nums[i] {
				STATE = UP
				maxLength++
			} else if nums[i-1] > nums[i] {
				STATE = DOWN
				maxLength++
			}
		case UP:
			if nums[i-1] > nums[i] {
				STATE = DOWN
				maxLength++
			}
		case DOWN:
			if nums[i-1] < nums[i] {
				STATE = UP
				maxLength++
			}
		}
	}
	return maxLength
}

func wiggleMaxLength1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	a, b := 1, 1
	for idx := 1; idx < length; idx++ {
		if nums[idx] > nums[idx-1] {
			a = b + 1
		}
		if nums[idx] < nums[idx-1] {
			b = a + 1
		}

	}
	if a > b {
		return a
	}
	return b
}

func wiggleMaxLength2(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	recordPN := make([]int, len(nums))
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] != 1 {
				if (nums[i]-nums[j])*recordPN[j] < 0 {
					if dp[i] < dp[j]+1 {
						dp[i] = dp[j] + 1
						if nums[i]-nums[j] > 0 {
							recordPN[i] = 1
						} else {
							recordPN[i] = -1
						}
						res = max(res, dp[i])
					}
				}
			} else {
				if nums[i]-nums[j] != 0 {
					if dp[i] < dp[j]+1 {
						dp[i] = dp[j] + 1
						res = max(res, dp[i])
						if nums[i]-nums[j] > 0 {
							recordPN[i] = 1
						} else {
							recordPN[i] = -1
						}
					}
				}
			}
		}
	}

	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
