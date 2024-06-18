package Solution

func Solution(numRows int) [][]int {
	pt := make([][]int, numRows)
	for i := range pt {
		pt[i] = make([]int, i+1)
	}
	for i, val := range pt {
		pt[i][0], pt[i][len(val)-1] = 1, 1
	}
	for i := 1; i < numRows-1; i++ {
		for j := 0; j < len(pt[i])-1; j++ {
			pt[i+1][j+1] = pt[i][j] + pt[i][j+1]
		}
	}
	return pt
}

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0] = 1
		if i >= 1 {
			dp[i][i] = 1
		}

	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}

	return dp
}
