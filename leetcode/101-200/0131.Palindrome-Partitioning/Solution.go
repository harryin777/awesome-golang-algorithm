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

func partition(s string) [][]string {
	res := make([][]string, 0, 10)
	var getPalindrome func(arr []string, nextIndex int)
	getPalindrome = func(arr []string, nextIndex int) {
		if nextIndex >= len(s) {
			res = append(res, arr)
			return
		}
		for i := nextIndex; i < len(s); i++ {
			currStr := s[nextIndex : i+1]
			if isPalindrome2(currStr) {
				arr = append(arr, currStr)
				getPalindrome(arr, i+1)
				arr = arr[0 : len(arr)-1]
			}
		}
	}
	getPalindrome([]string{}, 0)

	return res
}

func isPalindrome2(str string) bool {
	reversedStr := ""
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}

	return str == reversedStr
}

func partition3(s string) (ans [][]string) {
	n := len(s)

	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome2(s[i : j+1]) {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return
}
