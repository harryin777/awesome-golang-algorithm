package Solution

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(x, y int) bool {
		xStr := strconv.Itoa(nums[x])
		yStr := strconv.Itoa(nums[y])
		one, _ := strconv.Atoi(xStr + yStr)
		two, _ := strconv.Atoi(yStr + xStr)
		if one > two {
			return true
		} else {
			return false
		}
	})

	res := ""
	for _, val := range nums {
		res += strconv.Itoa(val)
	}

	return res
}
