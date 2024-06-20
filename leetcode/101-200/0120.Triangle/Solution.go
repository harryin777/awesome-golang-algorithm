package Solution

import (
	"fmt"
	"math"
	"strconv"
)

var minSum = 99999

// dfs solution
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	_dfs(triangle, 0, 0, "", 0)

	return minSum
}

func _dfs(triangle [][]int, i, j int, path string, sum int) int {
	//	terminator
	if i == len(triangle)-1 {
		sum += triangle[i][j]
		path += strconv.Itoa(triangle[i][j]) + " #"
		fmt.Println(path + "with sum: " + strconv.Itoa(sum))

		if sum < minSum {
			minSum = sum
			fmt.Println(minSum)
		}
		return sum
	}

	//	process
	sum += triangle[i][j]
	path += strconv.Itoa(triangle[i][j]) + " -> "

	//	drill down
	_dfs(triangle, i+1, j, path, sum)
	_dfs(triangle, i+1, j+1, path, sum)

	//	clear states
	return 0
}

// dp solution
func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			minSum := min(triangle[i+1][j], triangle[i+1][j+1])
			minSum += triangle[i][j]
			triangle[i][j] = minSum
		}
	}
	return triangle[0][0]
}

// dp solution
func minimumTotal2(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minimumTotal3(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	q := make([]int, len(triangle[0]))
	for i := 0; i < len(triangle[0]); i++ {
		q[i] += triangle[0][i]
		index := i
		for j := 1; j < len(triangle); j++ {
			if triangle[j][index] < triangle[j][index+1] {
				q[i] += triangle[j][index]
				index = index
			} else {
				q[i] += triangle[j][index+1]
				index = index + 1
			}
		}
	}

	res := math.MaxInt32
	for i := 0; i < len(q); i++ {
		res = min(res, q[i])
	}

	return res
}

func minimumTotal4(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[len(triangle)-1]))
	}
	dp[0][0] = triangle[0][0]
	for i := 0; i < len(triangle); i++ {
		for j := len(triangle[i]); j < len(triangle[len(triangle)-1]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}

		}
	}

	res := math.MaxInt32
	for i := 0; i < len(dp[len(dp)-1]); i++ {
		res = min(res, dp[len(dp)-1][i])
	}

	return res
}
