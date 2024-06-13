package Solution

import (
	"sort"
)

func groupAnagrams(ss []string) [][]string {
	tmp := make(map[int][]string, len(ss)/2)
	for _, s := range ss {
		c := encode(s)
		tmp[c] = append(tmp[c], s)
	}

	res := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		res = append(res, v)
	}

	return res
}

// prime 与 A～Z 对应，英文中出现概率越大的字母，选用的质数越小
var prime = []int{5, 71, 37, 29, 2, 53, 59, 19, 11, 83, 79, 31, 43, 13, 7, 67, 97, 23, 17, 3, 41, 73, 47, 89, 61, 101}

func encode(s string) int {
	res := 1
	for i := range s {
		res *= prime[s[i]-'a']
	}
	return res
}

func groupAnagrams2(strs []string) [][]string {
	strMap := make(map[string]string)
	emptyOne := make([]string, 0, len(strs))
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			emptyOne = append(emptyOne, "")
			continue
		}
		tmp := make([]byte, 0, len(strs[i]))
		for j := 0; j < len(strs[i]); j++ {
			tmp = append(tmp, strs[i][j])
		}
		sort.Slice(tmp, func(x, y int) bool {
			return tmp[x] < tmp[y]
		})
		sortedStr := ""
		for j := 0; j < len(tmp); j++ {
			sortedStr += string(tmp[j])
		}
		strMap[strs[i]] = sortedStr
	}

	res := make([][]string, 0, len(strs))
	fMap := make(map[string][]string)
	for key, val := range strMap {
		fMap[val] = append(fMap[val], key)
	}

	for _, val := range fMap {
		res = append(res, val)
	}
	if len(emptyOne) != 0 {
		res = append(res, emptyOne)
	}

	return res
}
