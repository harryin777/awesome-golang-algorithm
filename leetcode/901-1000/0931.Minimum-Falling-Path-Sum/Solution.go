package Solution

import "math"

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
