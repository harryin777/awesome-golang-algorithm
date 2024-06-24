package Solution

import "math"

func titleToNumber(columnTitle string) int {
	l := len(columnTitle)
	ans := 0
	for i := l - 1; i >= 0; i-- {
		ans += int(columnTitle[l-i-1]-'A'+1) * int(math.Pow(26, float64(i)))
	}

	return ans
}
