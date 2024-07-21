package Solution

import "strings"

func checkInclusion(s1 string, s2 string) bool {
	var dfs func(map[int]bool, string) bool
	dfs = func(indexDupMap map[int]bool, cur string) bool {
		if strings.Contains(s2, cur) {
			return true
		}
		if len(cur) >= len(s1) {
			return false
		}

		for i := 0; i < len(s1); i++ {
			if indexDupMap[i] {
				continue
			}
			indexDupMap[i] = true
			cur += string(s1[i])
			if dfs(indexDupMap, cur) {
				return true
			}
			delete(indexDupMap, i)
			cur = cur[:len(cur)-1]
		}

		return false
	}
	indexDupMap := make(map[int]bool)

	return dfs(indexDupMap, "")
}
