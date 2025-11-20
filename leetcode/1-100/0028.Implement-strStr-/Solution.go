package Solution

func strStr(haystack string, needle string) int {
	//	检查数据
	if len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for ; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}

		if j == len(needle) {
			return i
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	i, j, count := 0, 0, -1
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			count++
			i = count
			j = 0
		}
	}
	if i == len(haystack) && j != len(needle) {
		return -1
	}
	if j == len(needle) {
		return i - j
	}

	return -1
}

func strStr3(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	f := len(needle)
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		tmp := haystack[i : i+f]
		if tmp == needle {
			return i
		}
	}

	return -1
}
