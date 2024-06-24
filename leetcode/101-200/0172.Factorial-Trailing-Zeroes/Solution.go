package Solution

import "strconv"

func Solution(n int) int {
	count := 0
	countOfFive := make(map[int]int)
	for step := 5; step <= n; step += 5 {
		countOfFive[step] = 1
		if c, ok := countOfFive[step/5]; ok {
			countOfFive[step] += c
		}
		count += countOfFive[step]
	}
	return count
}

func Solution2(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n /= 5
	}
	return res
}

func trailingZeroes(n int) int {
	ans := n
	for n != 1 {
		ans *= n - 1
		n--
	}
	ansStr := strconv.FormatInt(int64(ans), 10)

	res := 0
	for i := len(ansStr) - 1; i >= 0; i-- {
		if ansStr[i] == '0' {
			res++
		} else {
			break
		}
	}

	return res
}
