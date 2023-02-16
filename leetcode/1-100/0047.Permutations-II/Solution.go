package Solution

import "sort"

func Solution(nums []int) [][]int {
	res := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	visited := make(map[int]bool)
	helper(nums, []int{}, visited, &res)
	return res
}

func helper(nums, res []int, visited map[int]bool, final *[][]int) {
	if len(nums) == len(res) {
		temp := make([]int, len(res))
		copy(temp, res)
		*final = append(*final, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		res = append(res, nums[i])
		visited[i] = true
		helper(nums, res, visited, final)
		visited[i] = false
		res = res[:len(res)-1]
	}
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var result [][]int
	var dfs func([]int, map[int]bool)
	dfs = func(path []int, indexMap map[int]bool) {
		if len(path) == len(nums) {
			data := make([]int, 0, len(nums))
			for _, v := range path {
				data = append(data, v)
			}
			result = append(result, data)
		}

		levelMap := make(map[int]bool)
		for i := 0; i < len(nums); i++ {
			if indexMap[i] || levelMap[nums[i]] {
				continue
			}
			levelMap[nums[i]] = true
			indexMap[i] = true
			path = append(path, nums[i])
			dfs(path, indexMap)
			indexMap[i] = false
			path = path[:len(path)-1]
		}
	}

	indexMap := make(map[int]bool)
	dfs([]int{}, indexMap)

	return result
}
