package Solution

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "*" || tokens[i] == "/" || tokens[i] == "+" || tokens[i] == "-" {
			num1, num2 := stack[0], stack[1]
			stack = stack[2:]
			var ans int
			switch tokens[i] {
			case "*":
				ans = num1 * num2
			case "/":
				ans = num2 / num1
			case "+":
				ans = num2 + num1
			case "-":
				ans = num2 - num1
			}
			stack = append([]int{ans}, stack...)
		} else {
			strInt, _ := strconv.ParseInt(tokens[i], 10, 64)
			stack = append([]int{int(strInt)}, stack...)
		}

	}

	return stack[0]
}
