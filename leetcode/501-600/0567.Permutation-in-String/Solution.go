package Solution

// checkInclusion 又超时了
func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 1
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}

	for right <= len(s2) {
		tmp := s2[left:right]
		_ = tmp
		for len(s1) != len(s2[left:right]) {
			right++
		}

		if isContain(s1, s2[left:right]) {
			return true
		}

		right++
		left++
	}

	return false
}

func isContain(s, t string) bool {
	var count1, count2 [26]int
	for index, val := range s {
		count1[val-'a']++
		count2[t[index]-'a']++
	}

	return count1 == count2
}
