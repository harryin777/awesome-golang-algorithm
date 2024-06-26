package Solution

import (
	"fmt"
	"strconv"
)

func Solution(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
	for i := 0; i < n; i++ {
		reversed := make([]int, n)
		for j := 0; j < n; j++ {
			reversed[j] = matrix[i][n-j-1]
		}
		matrix[i] = reversed
	}
}

func rotate(matrix [][]int) {
	tmp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		tmp[i] = make([]int, len(matrix[i]))
		copy(tmp[i], matrix[i])
	}

	n := len(matrix)
	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp[i]); j++ {
			matrix[j][n-i-1] = tmp[i][j]
		}
	}

}

func rotate180(matrix [][]int) {
	tmp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		tmp[i] = make([]int, len(matrix[i]))
		copy(tmp[i], matrix[i])
	}

	n := len(matrix)
	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp[i]); j++ {
			matrix[n-i-1][n-j-1] = tmp[i][j]
		}
	}

	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp[i]); j++ {
			fmt.Printf("%v,", strconv.Itoa(matrix[i][j]))
		}
		fmt.Println()
	}
}
