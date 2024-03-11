package Solution

import "strings"

func Solution(x bool) bool {
	return x
}

func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	var ans []int
	for right <= len(s) {
		for len(p) == len(s[left:right]) {
			right++
		}

		if check(p, s[left:right]) {
			ans = append(ans, left)
		}
		right++
		left++
	}

	return ans
}

func check(s, p string) bool {
	var cnt1, cnt2 [26]int
	for index, val := range s {
		cnt1[val-'a']++
		cnt2[p[index]-'a']++
	}

	return cnt1 == cnt2
}

func findAnagrams2(s string, p string) []int {
	if len(s) == 0 || len(s) < len(p) {
		return nil
	}
	res := make([]int, 0, 10)
	sumByte := 0
	for i := 0; i < len(p); i++ {
		sumByte += int(p[i])
	}
	left, right := 0, len(p)
	for right <= len(s) {
		if calSumByte(s[left:right]) == sumByte {
			if check2(left, right, p, s) {
				res = append(res, left)
			}
		}
		left++
		right++
	}

	return res
}

func calSumByte(s string) int {
	sumByte := 0
	for i := 0; i < len(s); i++ {
		sumByte += int(s[i])
	}
	return sumByte
}

func check2(left, right int, p, s string) bool {
	for i := left; i < right-left; i++ {
		if !strings.Contains(p, string(s[i])) {
			return false
		}
	}

	return true
}
