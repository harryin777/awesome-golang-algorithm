package Solution

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	var dfs func(int, int, int)
	dfs = func(x, y, pace int) {
		if len(res) == len(matrix)*len(matrix[0]) {
			return
		}

		if x < 0 || x >= m || y < 0 || y >= n || matrix[x][y] == -111 {
			preDir := pace % 4
			x, y = x-dir[preDir][0], y-dir[preDir][1]
			pace++
			nextDir := pace % 4
			nx, ny := x+dir[nextDir][0], y+dir[nextDir][1]
			dfs(nx, ny, pace)
			return
		}
		res = append(res, matrix[x][y])
		matrix[x][y] = -111
		nextDir := pace % 4
		nx, ny := x+dir[nextDir][0], y+dir[nextDir][1]
		dfs(nx, ny, pace)
	}

	dfs(0, 0, 0)

	return res
}
