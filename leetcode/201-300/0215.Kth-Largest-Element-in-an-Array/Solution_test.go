package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution func Info
type SolutionFuncType func([]int, int) int

var SolutionFuncList = []SolutionFuncType{
	findKthLargest_1,
	findKthLargest_2,
	findKthLargest,
}

// Test case info struct
type Case struct {
	name   string
	nums   []int
	k      int
	expect int
}

// Test case
var cases = []Case{
	{name: "TestCase 1", nums: []int{3, 2, 1, 5, 6, 4}, k: 2, expect: 5},
	{name: "TestCase 2", nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4, expect: 4},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.nums, c.k)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
