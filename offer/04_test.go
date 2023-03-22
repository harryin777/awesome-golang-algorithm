package offer

import (
	"fmt"
	"testing"
)

//

func TestFindNumberIn2DArray(t *testing.T) {
	fmt.Println(findNumberIn2DArray([][]int{
		{1, 3, 5, 7, 9}, {2, 4, 6, 8, 10}, {11, 13, 15, 17, 19}, {12, 14, 16, 18, 20}, {21, 22, 23, 24, 25},
	}, 13))
}

func TestExist(t *testing.T) {
	fmt.Println(exist([][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}, "AAB"))
}
