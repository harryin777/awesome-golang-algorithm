package Solution

import "math"

/*
	设状态为dp[i][j],表示A[0,i]与B[0,j]的最小编辑距离，对应形式分别为str1c,str2d
	如果c==d,f[i][j]=f[i-1][j-1]
	如果c!=d,
	a. 如果将c换成d,则f[i][j]=f[i-1][j-1]+1
	b. 如果在c后面加一个d,f[i][j]=f[i][j-1]+1
	c. 如果删掉c,f[i][j]=f[i-1][j]+1
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := [][]int{}
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
		}
	}
	// Print(dp)

	return dp[m][n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minDistance2(word1 string, word2 string) int {
	wl1 := len(word1) - 1
	wl2 := len(word2) - 1
	return dp(word1, wl1, word2, wl2)
}

// 把 1 转换成 2
// n 和 m 分别是两个字符串的指针,base case 是,如果1字符串已经走完,2字符串剩下的部分全部插入到 1,如果 2 已经走完,那么 1 需要把剩下的字符全部删除
// 所以为什么有 n+1 和 m+1
func dp(word1 string, n int, word2 string, m int) int {
	if n == -1 {
		return m + 1
	}
	if m == -1 {
		return n + 1
	}

	if word1[n] == word2[m] {
		return dp(word1, n-1, word2, m-1)
	}

	// 三种情况,分别是 插入, 删除, 替换
	return min3(dp(word1, n, word2, m-1)+1, dp(word1, n-1, word2, m)+1, dp(word1, n-1, word2, m-1)+1)
}

func min3(x, y, z int) int {
	return int(math.Min(math.Min(float64(x), float64(y)), float64(z)))
}

// memo 优化

var memo [][]int

func minDistance3(word1 string, word2 string) int {
	wl1 := len(word1)
	wl2 := len(word2)
	memo = make([][]int, wl1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, wl2)
		for j := 0; j < wl2; j++ {
			memo[i][j] = -1
		}
	}
	return dp2(word1, wl1-1, word2, wl2-1)
}

func dp2(word1 string, n int, word2 string, m int) int {
	if n == -1 {
		return m + 1
	}
	if m == -1 {
		return n + 1
	}

	if memo[n][m] != -1 {
		return memo[n][m]
	}

	if word1[n] == word2[m] {
		memo[n][m] = dp2(word1, n-1, word2, m-1)
		return memo[n][m]
	}

	// 三种情况,分别是 插入, 删除, 替换
	memo[n][m] = min3(dp2(word1, n, word2, m-1)+1, dp2(word1, n-1, word2, m)+1, dp2(word1, n-1, word2, m-1)+1)
	return memo[n][m]
}

func minDistance4(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}
func Min(args ...int) int {
	min := args[0]
	for _, item := range args {
		if item < min {
			min = item
		}
	}
	return min
}
