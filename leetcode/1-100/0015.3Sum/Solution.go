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
	res := make([][]int, 0, len(nums))
	var dfs func(int, []int, map[int]struct{})
	dfs = func(count int, currArr []int, visitedMap map[int]struct{}) {
		if count == 0 && len(currArr) == 3 {
			tmp := make([]int, len(currArr))
			copy(tmp, currArr)
			res = append(res, tmp)
			return
		}

		if len(visitedMap) == len(nums) {
			return
		}

		for i := 0; i < len(nums); i++ {
			if _, ok := visitedMap[nums[i]]; ok {
				continue
			}
			visitedMap[nums[i]] = struct{}{}
			currArr = append(currArr, nums[i])
			dfs(count+nums[i], currArr, visitedMap)
			delete(visitedMap, nums[i])
			currArr = currArr[0 : len(currArr)-1]
		}
	}
	visitedMap := make(map[int]struct{})
	dfs(0, []int{}, visitedMap)

	return res
}

// 这个解法可以，但是超出时间限制
func threeSum_3(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	removeDup := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := removeDup[nums[i]]; ok {
			continue
		}
		removeDup[nums[i]] = struct{}{}
		removeDup2 := make(map[int]struct{})
		for j := i + 1; j < len(nums); j++ {

			if _, ok := removeDup2[nums[j]]; ok {
				continue
			}
			removeDup2[nums[j]] = struct{}{}
			removeDup3 := make(map[int]struct{})
			for z := j + 1; z < len(nums); z++ {

				if _, ok := removeDup3[nums[z]]; ok {
					continue
				}
				removeDup3[nums[z]] = struct{}{}
				if nums[i]+nums[j]+nums[z] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[z]})
				}
			}
		}
	}

	return res
}

func threeSum_4(nums []int) [][]int {
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})
	res := make([][]int, 0, len(nums))
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			l, r := i+1, len(nums)-1
			for l < r {
				if nums[i]+nums[l]+nums[r] == 0 {
					res = append(res, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if nums[i]+nums[l]+nums[r] < 0 {
					l++
				} else {
					r--
				}
			}

		}

	}

	return res
}
