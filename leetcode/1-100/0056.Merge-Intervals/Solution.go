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
