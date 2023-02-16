package Solution

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	if candidates == nil || len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)
	set := []int{}
	explore(candidates, target, 0, &result, set)

	return result
}

func explore(candidates []int, target int, pos int, result *[][]int, set []int) {
	if target < 0 {
		return
	}

	if target == 0 {
		c := make([]int, len(set))
		copy(c, set)
		*result = append(*result, c)
		return
	}

	for i := pos; i < len(candidates); i++ {
		set = append(set, candidates[i])
		explore(candidates, target-candidates[i], i, result, set)
		set = set[:len(set)-1]
	}
}

func combinationSum3(candidates []int, target int) [][]int {
	var result [][]int
	var dfs func([]int, int, int, int)
	dfs = func(path []int, sum int, pos int, level int) {
		fmt.Printf("层 i : %v \n", level)
		if sum == 0 {
			data := make([]int, 0, len(path))
			for _, v := range path {
				data = append(data, v)
			}
			result = append(result, data)
			return
		}
		if sum < 0 {
			return
		}

		for i := pos; i < len(candidates); i++ {
			sum -= candidates[i]
			path = append(path, candidates[i])
			level++
			dfs(path, sum, i, level)
			level--
			sum += candidates[i]
			path = path[:len(path)-1]
		}
	}

	dfs([]int{}, target, 0, 1)

	return result
}
