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
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSum(&res, nums[i], nums[i+1:], target-nums[i])
	}
	return res
}

func threeSum(res *[][]int, first int, nums []int, target int) {
	nlen := len(nums)
	for i := 0; i < nlen-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, nlen-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				*res = append(*res, []int{first, nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}
}

// 解法可以,但是超出时间限制
func fourSum3(nums []int, target int) [][]int {
	n := len(nums)
	res := make([][]int, 0, n)
	sort.Ints(nums)
	var dfs func(int, int, []int)
	dfs = func(nextIndex, sum int, curArr []int) {
		if sum == target && len(curArr) == 4 {
			tmp := make([]int, len(curArr))
			copy(tmp, curArr)
			res = append(res, tmp)
			return
		}
		if len(curArr) >= 4 && sum != target {
			return
		}
		visitedMap := make(map[int]struct{})
		for i := nextIndex; i < len(nums); i++ {
			if _, ok := visitedMap[nums[i]]; ok {
				continue
			}
			visitedMap[nums[i]] = struct{}{}
			curArr = append(curArr, nums[i])
			dfs(i+1, sum+nums[i], curArr)
			curArr = curArr[0 : len(curArr)-1]
		}
	}

	dfs(0, 0, []int{})

	return res
}

func fourSum4(nums []int, target int) [][]int {
	l := len(nums)
	if l < 4 {
		return [][]int{}
	}

	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})

	res := make([][]int, 0, 10)
	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			n2, n3 := i+1, i+2
			n4 := l - 1
			for n2 < n4 && n3 < n4 {
				ans := nums[i] + nums[n2] + nums[n3] + nums[n4]
				if ans == target {
					res = append(res, []int{nums[i], nums[n2], nums[n3], nums[n4]})
					break
				} else if ans < target {
					if n3+1 < n4 {
						n3++
					} else {
						n2++
					}
				} else if ans > target {
					n4--
				}
			}

		}
	}

	return res
}

func fourSum5(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}
