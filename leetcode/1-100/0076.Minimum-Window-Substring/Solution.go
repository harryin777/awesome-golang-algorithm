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
