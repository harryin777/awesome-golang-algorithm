package Solution

import (
	"fmt"
	"testing"
)

func Test_minFallingPathSum(t *testing.T) {
	fmt.Println(minFallingPathSum([][]int{
		{100, -42, -46, -41},
		{31, 97, 10, -10},
		{-58, -51, 82, 89},
		{51, 81, 69, -51},
	}))
	fmt.Println(minFallingPathSum3([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))
}
