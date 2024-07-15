package Solution

var dirArr = [][]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func numIslands(grid [][]byte) int {
	res := 0
	queue := [][]int{}
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				visited[i][j] = true
				grid[i][j] = '0'
				queue = append(queue, []int{i, j})
				for len(queue) != 0 {
					cur := queue[0]
					queue = queue[1:]
					for k := 0; k < len(dirArr); k++ {
						x, y := cur[0]+dirArr[k][0], cur[1]+dirArr[k][1]
						if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || visited[x][y] {
							continue
						}
						if grid[x][y] == '1' {
							visited[x][y] = true
							queue = append(queue, []int{x, y})
							grid[x][y] = '0'
						}
					}
				}
			}
		}
	}

	return res
}
