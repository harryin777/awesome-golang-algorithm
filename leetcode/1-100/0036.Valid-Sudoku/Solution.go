package Solution

func isValidSudoku(board [][]byte) bool {
	n := len(board)
	rectSets := make([]int16, n*n/9) // sets for inner rectangles
	for i := 0; i < n; i++ {
		rowSet, colSet := int16(0), int16(0) // sets for a row and a column
		for j := 0; j < n; j++ {
			// checking a column and an inner rectangle
			if num := board[i][j]; num != '.' {
				numBit := int16(1 << (num - '0')) // calc a #num bit
				if rowSet&numBit > 0 {
					return false
				}
				rowSet |= numBit // set the bit

				// set the bit for a inner rectangle: ((n/3)*(i/3), j/3) - its coordinates
				if rectSets[(n/3)*(i/3)+j/3]&numBit > 0 {
					return false
				}
				rectSets[(n/3)*(i/3)+j/3] |= numBit
			}

			// checking column
			if num := board[j][i]; num != '.' {
				numBit := int16(1 << (num - '0'))
				if colSet&numBit > 0 {
					return false
				}
				colSet ^= numBit
			}
		}
	}
	return true
}

func isValidSudoku2(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		if checkRow(board, i) {
			for j := 0; j < len(board[i]); j++ {
				if i == 0 {
					if !checkColumn(board, j) {
						return false
					}
				}
				if i%3 == 0 && j%3 == 0 {
					if !checkBlock(i, j, board) {
						return false
					}
				}

			}
		} else {
			return false
		}
	}

	return true
}

func checkBlock(row, column int, board [][]byte) bool {
	c := make(map[byte]int)
	for i := row; i < row+3; i++ {
		for j := column; j < column+3; j++ {
			if board[i][j] != '.' {
				c[board[i][j]]++
			}
		}
	}
	for _, val := range c {
		if val > 1 {
			return false
		}
	}
	return true
}

func checkColumn(board [][]byte, index int) bool {
	c := make(map[byte]int)
	for i := 0; i < len(board); i++ {
		if board[i][index] != '.' {
			c[board[i][index]]++
		}

	}

	for _, val := range c {
		if val > 1 {
			return false
		}
	}

	return true
}

func checkRow(board [][]byte, index int) bool {
	c := make(map[byte]int)
	for i := 0; i < len(board); i++ {
		if board[index][i] != '.' {
			c[board[index][i]]++
		}
	}

	for _, val := range c {
		if val > 1 {
			return false
		}
	}

	return true
}
