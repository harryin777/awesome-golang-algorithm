package Solution

import "sort"

func Solution(x bool) bool {
	return x
}

/*
[1,2]
[2,3]
[2,4]
[3,4]
*/
func findLongestChain(pairs [][]int) int {

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0]-pairs[j][0] < 0
	})

	dp := make([]int, len(pairs))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(pairs); i++ {
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
