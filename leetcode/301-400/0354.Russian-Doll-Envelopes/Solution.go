package Solution

import (
	"fmt"
	"math"
	"sort"
)

func Solution(x bool) bool {
	return x
}

func maxEnvelopes(envelopes [][]int) int {
	sort.SliceStable(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1]-envelopes[j][1] > 0
		}

		return envelopes[i][0]-envelopes[j][0] < 0
	})
	fmt.Println(envelopes)
	dp := make([]int, len(envelopes))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}

	res := 0
	for i := 0; i < len(dp); i++ {
		res = int(math.Max(float64(res), float64(dp[i])))
	}
	return res
}
