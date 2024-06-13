package Solution

// 动态规划
func longestValidParentheses(s string) int {
	maxans := 0
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxans = Max(maxans, dp[i])
	}
	return maxans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Stack
func longestValidParentheses2(s string) int {
	maxans := 0
	stk := Stack{}

	stk.Push(-1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk.Push(i)
		} else {
			stk.Pop()
			if stk.IsEmpty() {
				stk.Push(i)
			} else {
				maxans = Max(maxans, i-stk.Top().(int))
			}
		}
	}
	return maxans
}

// 可以但是超出时间限制
func longestValidParentheses3(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			sub := s[i:j]
			if isValid(sub) {
				res = max(res, len(sub))
			}
		}
	}

	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func isValid(s string) bool {
	stack := make([]rune, len(s))
	top := 0
	for _, val := range s {
		if val == '(' {
			stack[top] = ')'
			top++
		} else if val == '{' {
			stack[top] = '}'
			top++
		} else if val == '[' {
			stack[top] = ']'
			top++
		} else {
			if top == 0 || stack[top-1] != val {
				return false
			}

			top--
		}
	}

	return top == 0
}

func longestValidParentheses4(s string) int {
	stack := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack += string(s[i])
		} else if s[i] == ')' {
			if len(stack) != 0 && stack[len(stack)-1] == '(' {
				stack = stack[0 : len(stack)-1]
			} else {
				stack += string(s[i])
			}
		}
	}

	return len(s) - len(stack)
}
