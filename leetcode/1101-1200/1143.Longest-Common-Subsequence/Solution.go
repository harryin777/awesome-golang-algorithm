package Solution

import (
	"fmt"
)

func Solution(x bool) bool {
	return x
}

// dp 在 index i,j 处的最长公共子序列,注意的是,dp 的长度一定是 text1 的长度+1,同理 text2
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func longestCommonSubsequence2(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

func longestCommonSubsequence3(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	if text1[0] == text2[0] {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}

	for i := 1; i < len(dp); i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < len(dp[0]); i++ {
		if text1[0] == text2[i] {
			dp[0][i] = 1
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
