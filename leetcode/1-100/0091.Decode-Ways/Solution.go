package Solution

func numDecodings(s string) int {
	// 思路：
	// 状态表示：f[i] i个数（最后一个数字是s[i-1]）解码得到的所有字符串的数量 f[i]从0开始
	// 状态计算： 集合可以分成：最后一个i是一位数，和最后一个i是两位数
	// 之前想过的错误思路：f[i] 表示以i结尾的decode共有多少种解法
	// 集合划分为不算i：f[i - j] + f[j] 和算i: f[i];
	n := len(s)
	// 加一是为了方便边界特判
	f := make([]int, n+1)
	// 空字符串特判
	f[0] = 1
	// 如果最后一个i是个位数 1 ~ 26
	// 难点：如何转化char为int
	for i := 1; i <= n; i++ { // 从一开始是因为表示的是前多少个字母，不是下标从一开始的字母
		// 如果最后一个i是个位数，且不为0
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		// 如果最后一个i是两位数，且在1~26之间
		if i >= 2 {
			t := (s[i-2]-'0')*10 + s[i-1] - '0'
			if t >= 10 && t <= 26 {
				f[i] += f[i-2]
			}
		}
	}
	return f[n]
}

func numDecodings2(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) && s[i-2] != '0' {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}
