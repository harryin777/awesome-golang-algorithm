package Solution

import "sort"

func Solution(x bool) bool {
	return x
}

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0, 10)
	var f func(startIndex int, sli []int)
	f = func(startIndex int, sli []int) {
		tmp := make([]int, len(sli))
		copy(tmp, sli)
		res = append(res, tmp)

		dup := make(map[int]struct{})
		for i := startIndex; i < len(nums); i++ {
			if _, ok := dup[nums[i]]; ok {
				continue
			}
			dup[nums[i]] = struct{}{}
			sli = append(sli, nums[i])
			f(i+1, sli)
			sli = sli[:len(sli)-1]
		}
	}
	sort.Ints(nums)
	sli := make([]int, 0, 10)
	f(0, sli)

	return res
}
