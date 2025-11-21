package Solution

import (
	"sort"
)

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

// 自定义排序规则
type SortByInt []Interval

func (p SortByInt) Len() int {
	return len(p)
}

func (p SortByInt) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p SortByInt) Less(i, j int) bool {
	return p[i].Start < p[j].Start
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals []Interval) []Interval {
	if intervals == nil || len(intervals) <= 1 {
		return intervals
	}
	//	先将结构体排序
	sort.Sort(SortByInt(intervals))

	start := intervals[0].Start
	end := intervals[0].End

	ans := make([]Interval, 0)

	for _, v := range intervals {
		if v.Start <= end {
			end = Max(end, v.End)
		} else {
			ans = append(ans, Interval{Start: start, End: end})
			start = v.Start
			end = v.End
		}
	}
	ans = append(ans, Interval{Start: start, End: end})
	return ans
}

func merge2(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(x, y int) bool {
		return intervals[x][1] < intervals[y][1]
	})
	res := make([][]int, 0, len(intervals))
	for len(intervals) > 1 {
		pre := intervals[0]
		last := intervals[1]
		i, j := pre[0], last[0]
		for i <= pre[1] {
			i++
		}
		i--
		if i < j {
			res = append(res, pre)
			intervals = intervals[1:]
			continue
		} else {
			newOne := []int{min(pre[0], last[0]), max(pre[1], last[1])}
			intervals = intervals[2:]
			intervals = append([][]int{newOne}, intervals...)
		}
	}
	res = append(res, intervals[0])

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func merge3(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(x, y int) bool {
		return intervals[x][0] < intervals[y][0]
	})
	i, j := 0, 1
	ans := make([][]int, 0, len(intervals))
	ansEmpty := true
	for j < len(intervals) {
		var slow []int
		if ansEmpty {
			slow = intervals[i]
		} else {
			slow = ans[len(ans)-1]
		}
		fast := intervals[j]
		if fast[0] <= slow[1] {
			newOne := make([]int, 2)
			if fast[1] <= slow[1] {
				newOne[1] = slow[1]
			} else {
				newOne[1] = fast[1]
			}
			newOne[0] = slow[0]
			if ansEmpty {
				ans = append(ans, newOne)
				intervals = intervals[2:]
			} else {
				ans = ans[:len(ans)-1]
				ans = append(ans, newOne)
				intervals = intervals[1:]
			}

		} else {
			if ansEmpty {
				ans = append(ans, slow)
			} else {
				ans = append(ans, fast)
			}
			intervals = intervals[1:]

		}
		j = 0
		ansEmpty = false
	}

	return ans
}

func merge4(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return false
		}
	})

	if len(intervals) == 0 {
		return nil
	}

	res := make([][]int, 0, 10)
	stack := make([][]int, 0, 10)
	stack = append(stack, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if stack[0][1] < intervals[i][1] {
			if stack[0][1] < intervals[i][0] {
				res = append(res, stack[0])
				stack = stack[:0]
				stack = append(stack, intervals[i])
				continue
			} else {
				stack[0][1] = intervals[i][1]
			}
		}

	}
	res = append(res, stack[0])
	return res
}
