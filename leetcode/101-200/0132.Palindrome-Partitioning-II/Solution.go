package Solution

func minCut(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if isPalindrome2(s[j:i]) {
				dp[i] = dp[j]
			} else {
				dp[i] = min(dp[i], dp[j]+1)
			}

		}
	}

	return dp[len(s)]
}

func isPalindrome2(str string) bool {
	reversedStr := ""
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}

	return str == reversedStr
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
