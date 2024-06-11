package Solution

import "fmt"

var xArr = []int{1, -1, 0, 0}
var yArr = []int{0, 0, 1, -1}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	res := 0
	outLine := make(map[string]bool)
	var dfs func(int, int, int)
	dfs = func(x, y, count int) {
		key := fmt.Sprintf("%v_%v", x, y)
		if _, ok := outLine[key]; ok {
			res++
			return
		}
		if x < 0 || x >= m || y < 0 || y >= n {
			res++
			outLine[key] = true
			return
		}
		if count >= maxMove {
			return
		}
		for i := 0; i < 4; i++ {
			nX := x + xArr[i]
			nY := y + yArr[i]
			dfs(nX, nY, count+1)

		}

	}

	dfs(startRow, startColumn, 0)

	return res
}
