package Solution

import "sort"

func combinationSum2(nums []int, target int) (result [][]int) {
	sort.Ints(nums)
	combinationSum2Helper(nums, nil, target, 0, 0, &result)
	return result
}

func combinationSum2Helper(nums, combo []int, target, sum, startIndex int, result *[][]int) {
	if sum == target {
		*result = append(*result, append([]int{}, combo...))
		return
	}
	for i := startIndex; i < len(nums) && (sum+nums[i]) <= target; i++ {
		if i != startIndex && nums[i] == nums[i-1] {
			continue
		}
		combinationSum2Helper(nums, append(combo, nums[i]), target, sum+nums[i], i+1, result)
	}
}

func combinationSum22(candidates []int, target int) [][]int {
	res := make([][]int, 0, 10)
	var dfs func([]int, int, int)
	dfs = func(arr []int, sum int, pos int) {
		if sum == target {
			res = append(res, append([]int(nil), arr...))
		}
		if sum > target {
			return
		}

		levelMap := make(map[int]struct{})
		for i := pos; i < len(candidates); i++ {
			if _, ok := levelMap[candidates[i]]; ok {
				continue
			}
			levelMap[candidates[i]] = struct{}{}
			arr = append(arr, candidates[i])
			dfs(arr, sum+candidates[i], i+1)
			arr = arr[:len(arr)-1]
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	dfs([]int{}, 0, 0)

	return res
}
