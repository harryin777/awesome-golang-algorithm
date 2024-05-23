package Solution

import (
	"fmt"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxF := make([]int, len(nums))
	minF := make([]int, len(nums))
	copy(maxF, nums)
	copy(minF, nums)

	for i := 1; i < len(nums); i++ {
		maxF[i] = max(maxF[i-1]*nums[i], max(nums[i], minF[i-1]*nums[i]))
		minF[i] = min(minF[i-1]*nums[i], min(nums[i], maxF[i-1]*nums[i]))
	}

	result := maxF[0]
	for _, v := range maxF {
		if v > result {
			result = v
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums)) // Output: 6
}
