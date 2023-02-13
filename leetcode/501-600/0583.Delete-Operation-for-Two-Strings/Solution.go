package Solution

func Solution(word1, word2 string) int {
	dp := make([][]int, 2)
	for idx := 0; idx < 2; idx++ {
		dp[idx] = make([]int, len(word2)+1)
	}
	for idx := 0; idx < len(word2)+1; idx++ {
		dp[0][idx] = idx
	}
	loop := 1
	for _, _byte := range []byte(word1) {
		dp[loop][0] = dp[1-loop][0] + 1
		for col := 1; col <= len(word2); col++ {
			leftTop := dp[1-loop][col-1]
			left := dp[loop][col-1]
			top := dp[1-loop][col]
			dp[loop][col] = leftTop
			if _byte != word2[col-1] {
				dp[loop][col] += 2
				if left+1 < dp[loop][col] {
					dp[loop][col] = left + 1
				}
				if top+1 < dp[loop][col] {
					dp[loop][col] = top + 1
				}
			}
		}

		loop = 1 - loop
	}
	return dp[1-loop][len(word2)]
}

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return 0
	}

	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return len(word2) - dp[len(word1)][len(word2)] + len(word1) - dp[len(word1)][len(word2)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
