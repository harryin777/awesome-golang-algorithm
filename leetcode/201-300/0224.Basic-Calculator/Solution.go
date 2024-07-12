package Solution

func calculate(s string) int {
	stack := []int{}
	preSign := '+'
	ans := 0
	for i := 0; i < len(s); i++ {
		var isDigit bool
		var num int
		if s[i] <= '9' && s[i] >= '0' {
			isDigit = true
		}
		if isDigit {
			num = num*10 + int(s[i]-'0')
		}
		if !isDigit && s[i] != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			num = 0
			preSign = int32(s[i])
		}
	}

	for _, val := range stack {
		ans += val
	}

	return ans
}

func calculate2(s string) int {
	res, _ := cal([]byte(s))
	return res
}

func cal(str []byte) (int, []byte) {
	preSign := '+'
	nums := 0
	stack := []int{}
	for len(str) > 0 {
		tmp := string(str)
		_ = tmp
		curC := str[0]
		str = str[1:]
		pre := string(preSign)
		_ = pre
		var isDigit bool
		if curC >= '0' && curC <= '9' {
			nums = nums*10 + int(curC-'0')
			isDigit = true
		}
		if !isDigit && curC != ' ' || len(str) == 0 {
			if curC == '(' {
				curNum, newStr := cal(str)
				nums = curNum
				str = newStr
			}
			switch preSign {
			case '*':
				stack[len(stack)-1] *= nums
			case '/':
				stack[len(stack)-1] /= nums
			case '+':
				stack = append(stack, nums)
			case '-':
				stack = append(stack, -nums)
			}
			nums = 0
			preSign = rune(curC)
		}
		if curC == ')' {
			break
		}

	}
	var ans int
	for _, val := range stack {
		ans += val
	}

	return ans, str
}
