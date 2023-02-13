package Solution

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs bool
		expect bool
	}{
		{"TestCase", true, true},
		{"TestCase", true, true},
		{"TestCase", false, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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

func TestMaxEnvelopes(t *testing.T) {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	fmt.Println(maxEnvelopes([][]int{{2, 100}, {3, 200}, {4, 300}, {5, 500}, {5, 400}, {5, 250}, {6, 370}, {6, 360}, {7, 380}}))
}

func TestSortSlice(t *testing.T) {
	slice := []int{8, 5, 3, 19, 22}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i]-slice[j] > 0
	})
	fmt.Println(slice)
}
