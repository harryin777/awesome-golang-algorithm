package Solution

import (
	"fmt"
)

func numDistinct(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func numDistinct3(s string, t string) int {
	l1, l2 := len(s), len(t)
	if l2 > l1 {
		return 0
	}
	dp := make([][]int, l1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = 1
	}
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[l1][l2]
}

func min(x int, y int, z int) int {
	if x > y {
		x = y
	}
	if x > z {
		x = z
	}
	return x
}

// 这个递归解法是正确的，但是时间过长
func numDistinct2(s string, t string) int {
	var res *int
	res = new(int)
	cal(s, t, "", 0, 0, res)

	return *res
}

func cal(s, t, tmp string, position, nextLevel int, res *int) {
	for i := nextLevel; i < len(s); i++ {
		if len(s[i:]) < len(t) && nextLevel == 0 {
			return
		}
		if position >= len(t) {
			return
		}
		if s[i] != t[position] {
			continue
		}
		tmp = fmt.Sprintf("%s%s", tmp, string(s[i]))
		if tmp == t {
			*res += 1
		}
		cal(s, t, tmp, position+1, i+1, res)
		tmp = tmp[0 : len(tmp)-1]
	}
}
