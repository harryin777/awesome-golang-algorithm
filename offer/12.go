package offer

//
//
type Character struct {
	Letter  byte
	Visited bool
}

func exist(board [][]byte, word string) bool {
	grid := make([][]Character, 0, len(board))
	for i := 0; i < len(board); i++ {
		tmp := make([]Character, 0, len(board[i]))
		for j := 0; j < len(board[i]); j++ {
			tmp = append(tmp, Character{
				Letter:  board[i][j],
				Visited: false,
			})
		}
		grid = append(grid, tmp)
	}
	depth := len(word)
	var dfs func(int, int, string, [][]Character, int) bool
	dfs = func(x int, y int, path string, grid [][]Character, curDepth int) bool {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || curDepth > depth || grid[x][y].Visited {
			return false
		}

		path += string(grid[x][y].Letter)
		if path[curDepth-1] != word[curDepth-1] {
			return false
		}

		grid[x][y].Visited = true

		if path == word {
			return true
		}

		if dfs(x, y+1, path, grid, curDepth+1) {
			return true
		}

		if dfs(x+1, y, path, grid, curDepth+1) {
			return true
		}

		if dfs(x-1, y, path, grid, curDepth+1) {
			return true
		}

		if dfs(x, y-1, path, grid, curDepth+1) {
			return true
		}
		grid[x][y].Visited = false
		path = path[:len(path)-1]

		return false

	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			newGrid := make([][]Character, len(grid))
			for m := 0; m < len(grid); m++ {
				newGrid[m] = make([]Character, len(grid[m]))
				copy(newGrid[m], grid[m])
			}
			if dfs(i, j, "", newGrid, 1) {
				return true
			}
		}
	}
	return false
}
