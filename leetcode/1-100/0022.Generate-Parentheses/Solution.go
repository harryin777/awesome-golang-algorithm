package Solution

func generateParenthesis(n int) []string {
	return dp("", 0, n*2)
}

func dp(prefix string, depth int, n int) []string {
	if depth == 0 && n == 0 {
		return []string{prefix}
	}
	var output []string
	if depth < n {
		output = append(output, dp(prefix+"(", depth+1, n-1)...)
	}
	if depth > 0 {
		output = append(output, dp(prefix+")", depth-1, n-1)...)
	}
	return output
}

func validCheck(str string, n int) bool {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			count++
		}
		if str[i] == ')' && count == 0 {
			return false
		} else if str[i] == ')' {
			count--
		}
	}

	return true
}

func generateParenthesis2(n int) []string {
	var res []string
	var dfs func(string, int, int)
	dfs = func(data string, open int, close int) {
		if len(data) == n*2 {
			res = append(res, data)
			return
		}

		if open < n {
			data += "("
			open++
			dfs(data, open, close)
			open--
			data = data[:len(data)-1]
		}

		if close < open {
			data += ")"
			close++
			dfs(data, open, close)
			close--
			data = data[:len(data)-1]
		}

	}

	dfs("", 0, 0)

	return res
}

var candidates = []string{
	"(",
	")",
}

func generateParenthesis3(n int) []string {
	count := n * 2
	ans := make([]string, 0, 10)
	var dfs func(string, int, int)
	dfs = func(data string, open, close int) {
		if len(data) == count {
			ans = append(ans, data)
			return
		}

		if open < n {
			data += "("
			dfs(data, open+1, close)
			data = data[:len(data)-1]
		}

		if close < open {
			data += ")"
			dfs(data, open, close+1)
			data = data[:len(data)-1]
		}
	}

	dfs("", 0, 0)

	return ans
}
