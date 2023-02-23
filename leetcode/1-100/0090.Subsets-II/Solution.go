package Solution

import "sort"

func Solution(x bool) bool {
	return x
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	var dfs func([]int, int, int, map[int]bool, []int)
	dfs = func(nums []int, depth int, pos int, used map[int]bool, data []int) {
		if depth == len(nums) {
			return
		}

		usedLevel1 := make(map[int]bool)
		for i := pos; i < len(nums); i++ {
			if used[i] || usedLevel1[nums[i]] {
				continue
			}

			usedLevel1[nums[i]] = true

			data = append(data, nums[i])
			used[i] = true
			data2 := make([]int, len(data))
			copy(data2, data)
			res = append(res, data2)
			dfs(nums, depth+1, i, used, data)
			data = data[:len(data)-1]
			used[i] = false
		}
		// usedLevel[nums[pos]] = false
	}

	used := make(map[int]bool)
	dfs(nums, 0, 0, used, []int{})
	res = append(res, []int{})

	return res
}
