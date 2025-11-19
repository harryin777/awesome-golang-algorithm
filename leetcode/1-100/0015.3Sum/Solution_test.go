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
type SolutionFuncType func([]int) [][]int

var SolutionFuncList = []SolutionFuncType{
	//threeSum_1,
	threeSum6,
}

// Test case info struct
type Case struct {
	name   string
	input  []int
	expect [][]int
}

// Test case
var cases = []Case{
	//{name: "TestCase 1", input: []int{}, expect: [][]int{}},
	//{name: "TestCase 2", input: []int{0}, expect: [][]int{}},
	{name: "TestCase 3", input: []int{-1, 0, 1, 2, -1, -4}, expect: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	//{name: "TestCase 4", input: []int{-2, 0, 0, 2, 2}, expect: [][]int{{-2, 0, 2}}},
	//{name: "TestCase 5", input: []int{1, -1, -1, 0}, expect: [][]int{{-1, 0, 1}}},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.input)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}

func Test_threeSum_2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args []int
		want [][]int
	}{
		// TODO: Add test cases.
		{name: "TestCase 3", args: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		//{name: "TestCase 4", args: []int{-2, 0, 0, 2, 2}, want: [][]int{{-2, 0, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, threeSum_4(tt.args), "threeSum_3(%v)", tt.args)
		})
	}
}
