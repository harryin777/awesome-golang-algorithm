package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// solution func Info
type SolutionFuncType func([]int, int) int

var SolutionFuncList = []SolutionFuncType{
	coinChange,
	coinChange2,
	coinChange3,
}

// test info struct
type Case struct {
	name   string
	coins  []int
	amount int
	expect int
}

// test case
var cases = []Case{
	{
		name:   "TestCase 1",
		coins:  []int{1, 2, 5},
		amount: 11,
		expect: 3,
	},
	{
		name:   "TestCase 2",
		coins:  []int{2},
		amount: 3,
		expect: -1,
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.coins, c.amount)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}

func TestCoinChange4(t *testing.T) {
	fmt.Println(CoinChange4([]int{1, 2, 5}, 11))
	fmt.Println(coinChange4([]int{1, 2, 5}, 11))
	fmt.Println(coinChange6([]int{1, 2, 5}, 11))

}
