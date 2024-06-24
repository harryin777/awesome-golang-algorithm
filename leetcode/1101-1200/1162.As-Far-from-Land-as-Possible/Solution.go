package Solution

import (
	"math"
)

var dirArr = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func maxDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxDis := math.MinInt32

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				queue := make([][]int, 0, 10)
				queue = append(queue, []int{i, j})
				visited := make([][]bool, m)
				for i := 0; i < m; i++ {
					visited[i] = make([]bool, n)
				}
				visited[i][j] = true

				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]

					for k := 0; k < len(dirArr); k++ {
						nx, ny := cur[0]+dirArr[k][0], cur[1]+dirArr[k][1]
						if nx < 0 || nx >= m || ny < 0 || ny >= n || visited[nx][ny] {
							continue
						}
						if grid[nx][ny] == 1 {
							maxDis = max(maxDis, abs(i-nx)+abs(j-ny))
							goto here
						}
						queue = append(queue, []int{nx, ny})
						visited[nx][ny] = true
					}
				}
			}
		here:
		}
	}

	if maxDis == math.MinInt32 {
		return -1
	}

	return maxDis
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
