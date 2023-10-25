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
