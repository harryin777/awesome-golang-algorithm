package Solution

func romanToInt(s string) int {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	sum := m[string(s[len(s)-1])]
	// 从后向前遍历
	// 每次和前面一位数比较
	for i := len(s) - 2; i >= 0; i-- {
		if m[string(s[i])] < m[string(s[i+1])] {
			sum -= m[string(s[i])]
		} else {
			sum += m[string(s[i])]
		}
	}

	return sum
}

func romanToInt2(s string) int {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	ans := 0
	// 从前向后遍历
	// 每次默认相加再检查和前面一位数的大小
	// 前面 > 后面  ans = s[i-1] + s[i]
	// 前面 < 后面  ans = - (2 * s[i-1]) + s[i] (因为默认加了一次所以需要减2次)
	for i := range s {
		ans += m[string(s[i])]
		if i > 0 && m[string(s[i])] > m[string(s[i-1])] {
			ans -= 2 * m[string(s[i-1])]
		}
	}
	return ans
}

func romanToInt3(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			res += 1000
		case 'C':
			if i+1 < len(s) && (s[i+1] == 'D' || s[i+1] == 'M') {
				res -= 100
			} else {
				res += 100
			}

		case 'D':
			res += 500
		case 'L':
			res += 50
		case 'X':
			if i+1 < len(s) && (s[i+1] == 'L' || s[i+1] == 'C') {
				res -= 10
			} else {
				res += 10
			}
		case 'V':
			res += 5
		case 'I':
			if i+1 < len(s) && (s[i+1] == 'V' || s[i+1] == 'X') {
				res -= 1
				break
			}

			res += 1
		}
	}

	return res
}
