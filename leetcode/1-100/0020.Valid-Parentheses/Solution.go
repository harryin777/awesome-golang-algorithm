package Solution

func isValid_1(s string) bool {
	stack := make([]rune, len(s))
	top := 0
	for _, v := range s {
		if v == '(' {
			stack[top] = ')'
			top++
		} else if v == '{' {
			stack[top] = '}'
			top++
		} else if v == '[' {
			stack[top] = ']'
			top++
		} else {
			if top == 0 || stack[top-1] != v {
				return false
			}
			top--
		}
	}
	return top == 0
}

func isValid_2(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func isValid2(s string) bool {
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

func isValid(s string) bool {
	stack := make([]byte, 0, 10)
	for _, val := range s {
		if val == '[' || val == '(' || val == '{' {
			stack = append([]byte{byte(val)}, stack...)
		} else if val == '}' {
			if len(stack) == 0 {
				return false
			}
			if stack[0] != '{' {
				return false
			}
			stack = stack[1:]
		} else if val == ')' {
			if len(stack) == 0 {
				return false
			}
			if stack[0] != '(' {
				return false
			}
			stack = stack[1:]
		} else if val == ']' {
			if len(stack) == 0 {
				return false
			}
			if stack[0] != '[' {
				return false
			}
			stack = stack[1:]
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
