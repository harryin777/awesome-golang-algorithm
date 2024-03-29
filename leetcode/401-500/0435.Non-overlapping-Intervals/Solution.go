package Solution

import "sort"

// 超出时间限制
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var slice []int
	for index, interval := range intervals {
		tmp := interval
		value := 1
		for j := index + 1; j < len(intervals); j++ {
			if tmp[1] <= intervals[j][0] {
				tmp = intervals[j]
				value++
			}
		}
		slice = append(slice, value)
	}

	sort.Ints(slice)
	return len(slice) - slice[len(slice)-1]

}

func eraseOverlapIntervals2(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	dp := make([]int, len(intervals))
	dp[0] = 1
	preMaxRight := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		curOne := intervals[i]
		if curOne[0] > preMaxRight {
			dp[i] = dp[i-1] + 1
			preMaxRight = curOne[1]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(intervals)-1]
}
