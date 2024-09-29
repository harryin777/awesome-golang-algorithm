package Solution

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	// dp[i] 代表以i结尾的最长有效括号长度
	dp := make([]int, len(s))
	ans := 0
	for i := 1; i < len(dp); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 1 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if s[i-1] == ')' {
				if i > dp[i-1] && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1] >= 2 {
						dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
			ans = max(ans, dp[i])
		}
	}
	return ans
}
func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
