package Solution

import "strings"

func findAllConcatenatedWordsInADict(words []string) []string {
	strContainsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if words[i] == words[j] {
				continue
			}
			if strings.Contains(words[i], words[j]) {
				strContainsMap[words[i]]++
			}
		}
	}

	res := make([]string, 0, len(words))
	for key, val := range strContainsMap {
		if val > 1 {
			res = append(res, key)
		}
	}

	return res
}
