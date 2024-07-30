package Solution

import "strings"

func wordBreak(s string, wordDict []string) []string {
	return dfs(s, wordDict, map[string][]string{})
}

func dfs(s string, wordDict []string, dict map[string][]string) []string {
	if res, ok := dict[s]; ok {
		return res
	}
	if len(s) == 0 {
		return []string{""}
	}

	var res []string

	for _, word := range wordDict {
		if len(word) <= len(s) && word == s[:len(word)] {
			for _, str := range dfs(s[len(word):], wordDict, dict) {
				if len(str) == 0 {
					res = append(res, word)
				} else {
					res = append(res, word+" "+str)
				}
			}
		}
	}
	dict[s] = res
	return res
}

func wordBreak2(s string, wordDict []string) []string {
	res := make([]string, 0, 10)
	var dfs func(c int, words []string)
	dfs = func(c int, words []string) {
		if c == len(s) {
			tmp := strings.Join(words, " ")
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(wordDict); i++ {
			if c+len(wordDict[i]) <= len(s) && s[c:c+len(wordDict[i])] == wordDict[i] {
				words = append(words, wordDict[i])
				dfs(c+len(wordDict[i]), words)
				words = words[:len(words)-1]
			}
		}
	}
	dfs(0, []string{})

	return res
}

func wordBreak3(s string, wordDict []string) []string {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	ans := make([]string, 0, 10)
	var dfs func(string, []string)
	dfs = func(curStr string, arr []string) {
		if len(curStr) == 0 {
			ansStr := strings.Join(arr, " ")
			ansStr = strings.TrimSuffix(ansStr, " ")
			ans = append(ans, ansStr)
			return
		}

		for i := 1; i <= len(curStr); i++ {
			str := curStr[:i]
			if wordMap[str] {
				newStr := curStr[:i]
				arr = append(arr, newStr)
				dfs(curStr[i:], arr)
				arr = arr[:len(arr)-1]
			}
		}
	}

	dfs(s, []string{})

	return ans
}
