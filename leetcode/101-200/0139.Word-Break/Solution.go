package Solution

func canBreak(start int, s string, wordMap map[string]bool) bool {
	if start == len(s) {
		return true
	}
	for i := start + 1; i <= len(s); i++ {
		prefix := s[start:i]
		if wordMap[prefix] && canBreak(i, s, wordMap) {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	return canBreak(0, s, wordMap)
}

func canBreak2(start int, s string, wordMap map[string]bool, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}
	if res, ok := memo[start]; ok {
		return res
	}
	for i := start + 1; i <= len(s); i++ {
		prefix := s[start:i]
		if wordMap[prefix] && canBreak2(i, s, wordMap, memo) {
			memo[start] = true
			return true
		}
	}
	memo[start] = false
	return false
}

func wordBreak2(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	memo := make(map[int]bool)
	return canBreak2(0, s, wordMap, memo)
}

func wordBreakBFS(s string, wordDict []string) bool {
	l := len(s)
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	queue := []int{}
	queue = append(queue, 0)
	visited := map[int]bool{}

	for len(queue) != 0 {
		start := queue[0]
		queue = queue[1:]
		if visited[start] {
			continue
		}
		visited[start] = true
		for i := start + 1; i <= l; i++ {
			prefix := s[start:i]
			if wordMap[prefix] {
				if i < l {
					queue = append(queue, i)
				} else {
					return true
				}
			}
		}
	}
	return false
}

func wordBreak3(s string, wordDict []string) bool {
	wordMap := make(map[string]string)
	for _, s2 := range wordDict {
		wordMap[s2] = s2
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if _, ok := wordMap[s[i:j]]; ok {
				dp[j] = dp[j] || dp[i]
			}
		}
	}

	return dp[len(s)]
}

func wordBreak4(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func wordBreak5(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	dp[0] = false
	j := 0
	for i := 0; i <= len(s); i++ {
		for k := 0; k < len(wordDict); k++ {
			if s[j:i] == wordDict[k] {
				if j == 0 {
					dp[i] = true
				} else {
					dp[i] = dp[j] && true
				}
				j = i
				break
			}
		}

	}

	return dp[len(s)-1]
}

func wordBreak6(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			sub := s[j:i]
			dp[i] = dp[i] || (dp[j] && wordMap[sub])
		}
	}

	return dp[len(s)]
}

func wordBreak7(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			sub := s[j:i]
			if wordMap[sub] {
				dp[i] = dp[i] || dp[i-len(sub)]
			}
		}
	}

	return dp[len(s)]
}
