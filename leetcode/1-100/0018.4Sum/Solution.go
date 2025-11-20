package Solution

import (
	"sort"
)

func fourSum(nums []int, target2 int) [][]int {
	sort.Ints(nums)
	var re [][]int
	for i1 := 0; i1 < len(nums); i1++ {
		for i := i1 + 1; i < len(nums); i++ {
			target := -nums[i] - nums[i1] + target2
			front := i + 1
			back := len(nums) - 1
			for front < back {
				sum := nums[front] + nums[back]
				if sum < target {
					front++
				} else if sum > target {
					back--
				} else {
					resub := make([]int, 4)
					resub[0] = nums[i]
					resub[1] = nums[front]
					resub[2] = nums[back]
					resub[3] = nums[i1]
					re = append(re, resub)
					for front < back && nums[front] == resub[1] {
						front++
					}
					for back > front && nums[back] == resub[2] {
						back--
					}
				}
			}
			for i+1 < len(nums) && nums[i+1] == nums[i] {
				i++
			}
		}
		for i1+1 < len(nums) && nums[i1+1] == nums[i1] {
			i1++
		}
	}
	return re
}

func fourSum1(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			dup := make(map[int]struct{})
			for l < r {
				val := nums[i] + nums[j] + nums[l] + nums[r]
				if val == target {
					if _, ok := dup[nums[l]]; ok {
						l++
						continue
					}
					dup[nums[l]] = struct{}{}
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				} else if val < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return res
}
