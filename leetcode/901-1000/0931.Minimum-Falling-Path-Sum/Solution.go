package Solution

import (
	"math"
)

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n := 0, len(matrix[0])-1

	return dp(m, n, 0, 0, matrix)

}

func dp(left, right, sum, index int, matrix [][]int) int {
	if index >= len(matrix) {
		return sum
	}

	res := math.MaxInt32
	var newLeft, newRight int
	for i := left; i <= right; i++ {
		newLeft = i - 1
		newRight = i + 1
		if newLeft < 0 {
			newLeft = 0
		}
		if newRight >= len(matrix[0]) {
			newRight = len(matrix[0]) - 1
		}
		res = matrix[index][i]
		res = min(dp(newLeft, newRight, sum+res, index+1, matrix), res) + sum
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// 非常朴素的一种方法,直接每行每列遍历,然后当前位置的最小值就是上一行相邻三个位置的最小值加上当前位置的值
func minFallingPathSum2(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	copy(dp[0], matrix[0])
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			mn := dp[i-1][j]
			if j > 0 {
				mn = min(mn, dp[i-1][j-1])
			}
			if j < n-1 {
				mn = min(mn, dp[i-1][j+1])
			}
			dp[i][j] = mn + matrix[i][j]
		}
	}
	minVal := dp[n-1][0]
	for _, val := range dp[n-1] {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func minFallingPathSum3(matrix [][]int) int {
	var dp [][]int
	rowCount := len(matrix)
	dp = make([][]int, rowCount)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = matrix[0][i]
	}

	for i := 1; i < rowCount; i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = cal(dp[i-1], j-1, j+1) + matrix[i][j]
		}
	}

	res := 99999
	for i := 0; i < len(dp[rowCount-1]); i++ {
		res = min(res, dp[rowCount-1][i])
	}

	return res
}

func cal(lastRow []int, begin, end int) int {
	if begin < 0 {
		begin = 0
	}
	if end >= len(lastRow) {
		end = len(lastRow) - 1
	}

	res := 99999
	for i := begin; i <= end; i++ {
		res = min(res, lastRow[i])
	}

	return res
}
