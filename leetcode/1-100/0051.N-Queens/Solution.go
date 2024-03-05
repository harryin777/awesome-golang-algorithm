package Solution

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(x int) [][]string {
	board := make([][]string, x)
	for i := range board {
		board[i] = make([]string, x)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	res := make([][]string, 0)
	helper(board, &res, 0)
	return res
}

func helper(board [][]string, res *[][]string, row int) {
	if len(board) == row {
		temp := make([]string, 0)
		for i := range board {
			temp = append(temp, strings.Join(board[i], ""))
		}
		*res = append(*res, temp)
		return
	}
	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			board[row][col] = "Q"
			helper(board, res, row+1)
			board[row][col] = "."
		}
	}
}

func isValid(board [][]string, r, c int) bool {
	for i := 0; i < r; i++ {
		if board[i][c] == "Q" {
			return false
		}
	}
	for i, j := r, c; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	for i, j := r, c; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0, 10)
	var str string
	for i := 0; i < n; i++ {
		str = str + "."
	}
	var dfs func(c []string, used map[string]struct{}, row int)
	dfs = func(c []string, used map[string]struct{}, row int) {
		for i := 0; i < n; i++ {
			pos := fmt.Sprintf("%v:%v", row, i)

			if !checkUnusedMap(row, i, used) {
				continue
			}
			str1 := []byte(str)
			str1[i] = 'Q'
			str2 := string(str1)
			c = append(c, str2)
			used[pos] = struct{}{}
			dfs(c, used, row+1)
			delete(used, pos)
			c = c[:len(c)-1]
		}

		if len(c) == n {
			tmp := make([]string, len(c))
			copy(tmp, c)
			res = append(res, tmp)
			return
		}
	}
	used := make(map[string]struct{})
	dfs([]string{}, used, 0)

	return res
}

func checkUnusedMap(x, y int, used map[string]struct{}) bool {

	for val, _ := range used {
		arrs := strings.Split(val, ":")
		x1, _ := strconv.ParseInt(arrs[0], 10, 64)
		y1, _ := strconv.ParseInt(arrs[1], 10, 64)
		if (x1-y1) == int64(x-y) || (x1+y1) == int64(x+y) {
			return false
		}
		if x1 == int64(x) || y1 == int64(y) {
			return false
		}
	}

	return true
}
