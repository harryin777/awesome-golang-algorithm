package Solution

import (
	"fmt"
	"strings"
)

func minWindow(s string, t string) string {
	hash := make(map[uint8]int)
	cut := 0

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		if _, ok := hash[s[i]]; !ok {
			cut++
		}
		hash[s[i]]++
	}

	res := ""
	for i, j, c := 0, 0, 0; i < len(s); i++ {
		if hash[s[i]] == 1 {
			c++
		}
		hash[s[i]]--
		for c == cut && hash[s[i]] < 0 {
			hash[s[j]]++
			j++
		}
		if c == cut {
			if res == "" || len(res) > i-j+1 {
				res = s[j : i-j+1]
			}
		}

	}
	return res
}

// 这个解法是可以的，但是超时了
func minWindow2(s string, t string) string {
	left, right := 0, 1
	if !isContain(s, t) {
		return ""
	}

	res := s

	for right < len(s) {
		tmp := s[left:right]
		if !isContain(tmp, t) {
			right++
		}

		tmp = s[left:right]
		for isContain(tmp, t) {
			if len(res) > right-left {
				res = tmp
			}
			left++
			tmp = s[left:right]
		}
	}

	return res
}

func isContain(s, t string) bool {
	//list := make([]int, 256)
	for _, val := range t {
		//list[val]++
		if strings.Count(s, string(val)) < strings.Count(t, string(val)) {
			return false
		}
	}

	//for _, val := range s {
	//	if strings.Contains(t, string(val)) {
	//		list[val]--
	//	}
	//}
	//
	//for _, val := range list {
	//	if val > 0 {
	//		return false
	//	}
	//}

	return true && len(s) >= len(t)
}

func minWindow3(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	res := s + t
	var tCount [58]int
	for i := 0; i < len(t); i++ {
		tCount[t[i]-'A']++
	}
	left, right := 0, len(t)
	for right <= len(s) {
		if !check(s[left:right], t, tCount) {
			right++
			continue
		}
		res = min(res, s[left:right])
		left++
		for check(s[left:right], t, tCount) {
			res = min(res, s[left:right])
			left++
		}
	}

	if res == s+t {
		return ""
	}

	return res
}

func min(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s2
	}

	return s1
}

func check(str1, str2 string, tCount [58]int) bool {
	var s1 [58]int
	for i := 0; i < len(str1); i++ {
		s1[str1[i]-'A']++
	}
	for i := 0; i < len(str2); i++ {
		if s1[str2[i]-'A'] < tCount[str2[i]-'A'] {
			return false
		}
	}

	return true
}

func minWindow4(s string, t string) string {
	if len(s) < len(t) || len(s) == 0 {
		return ""
	}
	lowerFlag := false
	if s[0] < 'a' {
		s = strings.ToLower(s)
		t = strings.ToLower(t)
		lowerFlag = true
	}

	count := make([][][]int, len(s))
	for i := 0; i < len(s); i++ {
		count[i] = make([][]int, len(s)+1)
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			count[i][j] = make([]int, 26)
			for _, val := range s[i:j] {
				count[i][j][val-'a']++
			}
		}
	}

	res := s + t
	var tCount [26]int
	for i := 0; i < len(t); i++ {
		tCount[t[i]-'a']++
	}
	left, right := 0, len(t)
	for left < right && right <= len(s) {
		if !check2(count[left][right], t, tCount) {
			right++
			continue
		}
		res = min(res, s[left:right])
		left++
		for left < right && check2(count[left][right], t, tCount) {
			res = min(res, s[left:right])
			left++
		}
	}

	if res == s+t {
		return ""
	}

	if lowerFlag {
		res = strings.ToUpper(res)
	}

	return res
}

func check2(sCount []int, t string, tCount [26]int) bool {
	for i := 0; i < len(t); i++ {
		if sCount[t[i]-'a'] < tCount[t[i]-'a'] {
			return false
		}
	}

	return true
}
