package Solution

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	ans := nums[0] + nums[1] + nums[len(nums)-1]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		start, end := i+1, len(nums)-1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum > target {
				end--
			} else {
				start++
			}

			if Abs(sum-target) < Abs(ans-target) {
				ans = sum
			}

		}
	}
	return ans
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func threeSumClosest2(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})
	res := nums[0] + nums[1] + nums[len(nums)-1]
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < target {
				l++
			} else { // 注意这里不能是 else if 会出现一直不大于也不小于循环出不来的情况
				r--
			}
			if Abs(sum-target) < Abs(res-target) {
				res = sum
			}
		}

	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
