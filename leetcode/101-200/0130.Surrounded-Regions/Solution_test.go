package Solution

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]byte
		expect bool
	}{
		{
			name: "t1",
			inputs: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expect: false,
		},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			solve(c.inputs)
			for i := 0; i < len(c.inputs); i++ {
				for j := 0; j < len(c.inputs[i]); j++ {
					fmt.Print(string(c.inputs[i][j]))
				}
				fmt.Println()
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
