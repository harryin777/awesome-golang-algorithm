package Solution

func longestCommonPrefix(strs []string) string {
	var result string

	maxIndex := 0
	for i, str := range strs {
		strLen := len(str)
		if i == 0 {
			maxIndex = strLen - 1
			result = str
			continue
		}

		if strLen-1 < maxIndex {
			maxIndex = strLen - 1
			result = result[:strLen]
		}

		for j := 0; j <= maxIndex && j < strLen; j++ {
			if str[j] != result[j] {
				maxIndex = j - 1
				result = str[:j]
			}
		}
	}

	return result
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for i := 1; i < len(strs); i++ {
		len := min(len(res), len(strs[i]))
		j := 0
		tmp := ""
		for ; j < len; j++ {
			if res[j] == strs[i][j] {
				tmp += string(res[j])
				continue
			} else {
				break
			}
		}
		if j == 0 {
			res = ""
			break
		}

		res = tmp
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
