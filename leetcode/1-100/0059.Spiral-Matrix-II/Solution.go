package Solution

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func generateMatrix(n int) [][]int {
	maxOne := n * n
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	var dfs func(int, int, int, int)
	dfs = func(x, y, cur, pace int) {
		if cur > maxOne {
			return
		}

		if x < 0 || x >= n || y < 0 || y >= n || matrix[x][y] != 0 {
			preDir := pace % 4
			x, y = x-dir[preDir][0], y-dir[preDir][1]
			pace++
			nextDir := pace % 4
			nx, ny := x+dir[nextDir][0], y+dir[nextDir][1]
			dfs(nx, ny, cur, pace)
			return
		}
		matrix[x][y] = cur
		cur++
		nextDir := pace % 4
		nx, ny := x+dir[nextDir][0], y+dir[nextDir][1]
		dfs(nx, ny, cur, pace)
	}
	dfs(0, 0, 1, 0)

	return matrix
}
