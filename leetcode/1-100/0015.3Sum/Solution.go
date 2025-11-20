package Solution

import (
	"sort"
)

func threeSum_1(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) <= 2 {
		return ans
	}

	//	排序
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			low, high, sum := i+1, len(nums)-1, 0-nums[i]
			for low < high {
				//	和为0的时候
				if nums[low]+nums[high] == sum {
					ans = append(ans, []int{nums[i], nums[low], nums[high]})
					//	排除；连续相等的数
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if nums[low]+nums[high] < sum {
					low++
				} else {
					high--
				}
			}
		}
	}

	return ans
}

func threeSum_2(nums []int) [][]int {
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})
	res := make([][]int, 0, 0)
	dup := make(map[int]struct{})
	for i := 0; i < len(nums)-2; i++ {
		if _, ok := dup[nums[i]]; ok {
			continue
		}
		dup[nums[i]] = struct{}{}
		lDup := make(map[int]struct{})
		l, r := i+1, len(nums)-1
		for l < r {
			val := nums[i] + nums[l] + nums[r]
			if val < 0 {
				l++
			} else if val == 0 {
				if _, ok := lDup[nums[l]]; ok {
					l++
					continue
				}
				lDup[nums[l]] = struct{}{}
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else {
				r--
			}
		}
	}

	return res
}
