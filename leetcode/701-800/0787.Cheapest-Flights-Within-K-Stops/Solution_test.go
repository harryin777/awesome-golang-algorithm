package Solution

import (
	"fmt"
	"testing"
)

// 使用案列
func ExampleSolution() {
}

func Test_t(t *testing.T) {
	fmt.Println(findCheapestPrice(3, [][]int{{1, 2, 100}, {0, 1, 100}, {0, 2, 500}}, 0, 2, 1))
}
