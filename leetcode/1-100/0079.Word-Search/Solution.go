package Solution

import "fmt"

func exist(borad [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}
	b := []byte(word)
	for x := range borad {
		for y := range borad[x] {
			if dfs(x, y, borad, b) {
				return true
			}
		}
	}
	return false
}

func dfs(x, y int, board [][]byte, word []byte) bool {
	if x < 0 || x >= len(board) {
		return false
	}
	if y < 0 || y >= len(board[x]) {
		return false
	}
	if board[x][y] != word[0] {
		return false
	}
	if len(word) == 1 {
		return true
	}
	tmp := board[x][y]
	board[x][y] = '-'
	found := dfs(x+1, y, board, word[1:]) ||
		dfs(x-1, y, board, word[1:]) ||
		dfs(x, y+1, board, word[1:]) ||
		dfs(x, y-1, board, word[1:])
	board[x][y] = tmp
	return found
}

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// 解法可以，但是超时间
func exist2(board [][]byte, word string) bool {
	l := len(word)
	var dfs func(int, int, int, string, map[string]struct{}) bool
	dfs = func(x, y, index int, cur string, visited map[string]struct{}) bool {
		if len(cur) == l {
			return true
		}

		if board[x][y] != word[index] {
			return false
		}

		visited[fmt.Sprintf("%v%v", x, y)] = struct{}{}
		cur += string(board[x][y])
		for _, dir := range dirs {
			if x+dir[0] < 0 || x+dir[0] >= len(board) || y+dir[1] < 0 || y+dir[1] >= len(board[0]) {
				continue
			}
			if _, ok := visited[fmt.Sprintf("%v%v", x+dir[0], y+dir[1])]; ok {
				continue
			}
			if dfs(x+dir[0], y+dir[1], index+1, cur, visited) {
				return true
			}
		}
		delete(visited, fmt.Sprintf("%v%v", x, y))
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != word[0] {
				continue
			}
			visited := make(map[string]struct{})
			if dfs(i, j, 0, "", visited) {
				return true
			}
		}
	}

	return false
}
