package Solution

import (
	"strconv"
	"unicode"
)

func diffWaysToCompute(input string) []int {
	return ways(input, map[string][]int{})
}

func calculate(a, b int, operate rune) int {
	switch operate {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	default:
		panic("operate not exist")
	}
}

func ways(input string, cache map[string][]int) []int {
	if _, ok := cache[input]; ok {
		return cache[input]
	}

	ans := []int{}

	for i := 0; i < len(input); i++ {
		ch := input[i]

		if ch == '+' || ch == '-' || ch == '*' {
			left := input[:i]
			right := input[i+1:]

			l := ways(left, cache)
			r := ways(right, cache)

			for _, a := range l {
				for _, b := range r {
					ans = append(ans, calculate(a, b, rune(ch)))
				}
			}
		}
	}

	if len(ans) == 0 {
		number, _ := strconv.Atoi(input)
		ans = append(ans, number)
	}

	cache[input] = ans

	return ans
}

const addition, subtraction, multiplication = -1, -2, -3

func diffWaysToCompute2(expression string) []int {
	ops := []int{}
	for i, n := 0, len(expression); i < n; {
		if unicode.IsDigit(rune(expression[i])) {
			x := 0
			for ; i < n && unicode.IsDigit(rune(expression[i])); i++ {
				x = x*10 + int(expression[i]-'0')
			}
			ops = append(ops, x)
		} else {
			if expression[i] == '+' {
				ops = append(ops, addition)
			} else if expression[i] == '-' {
				ops = append(ops, subtraction)
			} else {
				ops = append(ops, multiplication)
			}
			i++
		}
	}

	n := len(ops)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
	}
	var dfs func(int, int) []int
	dfs = func(l, r int) []int {
		res := dp[l][r]
		if res != nil {
			return res
		}
		if l == r {
			dp[l][r] = []int{ops[l]}
			return dp[l][r]
		}
		for i := l; i < r; i += 2 {
			left := dfs(l, i)
			right := dfs(i+2, r)
			for _, x := range left {
				for _, y := range right {
					if ops[i+1] == addition {
						dp[l][r] = append(dp[l][r], x+y)
					} else if ops[i+1] == subtraction {
						dp[l][r] = append(dp[l][r], x-y)
					} else {
						dp[l][r] = append(dp[l][r], x*y)
					}
				}
			}
		}
		return dp[l][r]
	}
	return dfs(0, n-1)
}
