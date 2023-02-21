package Solution

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
