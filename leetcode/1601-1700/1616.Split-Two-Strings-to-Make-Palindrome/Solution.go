package Solution

func checkPalindromeFormation(a string, b string) bool {
	if len(a) == len(b) && len(a) == 1 {
		return true
	}

	lower := len(a)
	if len(a) > len(b) {
		lower = len(b)
	}

	var index int

	for index <= lower {
		if isH(a[:index] + b[index:]) {
			return true
		}
		if isH(b[index:] + a[:index]) {
			return true
		}
		if isH(a[index:] + b[:index]) {
			return true
		}
		if isH(b[:index] + a[index:]) {
			return true
		}
		index++
	}

	return false
}

func isH(str string) bool {
	left, right := 0, len(str)-1
	for left < right {
		if str[left] == str[right] {
			left++
			right--
		} else {
			return false
		}
	}

	if len(str)%2 == 0 {
		return true
	}

	if str[left] == str[right] {
		return true
	}

	return false
}

func checkPalindromeFormation2(a, b string) bool {
	return checkConcatenation(a, b) || checkConcatenation(b, a)
}

func checkConcatenation(a, b string) bool {
	left, right := 0, len(a)-1
	for left < right && a[left] == b[right] {
		left++
		right--
	}
	if left >= right {
		return true
	}
	return checkSelfPalindrome(a[left:right+1]) || checkSelfPalindrome(b[left:right+1])
}

func checkSelfPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		left++
		right--
	}
	return left >= right
}
