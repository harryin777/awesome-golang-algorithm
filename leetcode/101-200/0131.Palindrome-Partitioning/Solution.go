package Solution

func Solution(s string) [][]string {
	res := make([][]string, 0)
	buildPartition(s, 0, []string{}, &res)
	return res
}

func buildPartition(s string, start int, path []string, res *[][]string) {
	if start >= len(s) {
		dst := make([]string, len(path))
		copy(dst, path)
		*res = append(*res, dst)
		return
	}

	for end := start + 1; end <= len(s); end++ {
		str := s[start:end]
		if isPalindrome(str) {
			buildPartition(s, end, append(path, str), res)
		}
	}
}

func isPalindrome(s string) bool {
	for _s, e := 0, len(s)-1; _s < e; _s, e = _s+1, e-1 {
		if s[_s] != s[e] {
			return false
		}
	}
	return true
}

func partition2(s string) [][]string {
	res := make([][]string, 0, 10)
	var getPalindrome func(str string, arr []string)
	getPalindrome = func(str string, arr []string) {
		if len(str) == 0 {
			res = append(res, arr)
			return
		}
		currArr := make([]string, len(arr))
		copy(currArr, arr)
		for i := 0; i < len(str); i++ {
			currStr := str[:i+1]
			if isPalindrome2(currStr) {
				currArr = append(currArr, currStr)
				getPalindrome(str[i+1:], currArr)
			}

		}
	}
	getPalindrome(s, []string{})

	return res
}

func isPalindrome2(str string) bool {
	reversedStr := ""
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}

	return str == reversedStr
}
