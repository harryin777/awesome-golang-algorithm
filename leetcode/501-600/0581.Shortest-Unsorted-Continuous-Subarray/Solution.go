package Solution

import (
	"math"
	"sort"
)

func Solution(nums []int) int {
	start, end, min, max := 0, -1, math.MaxInt32, math.MinInt32
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}
		if nums[j] <= min {
			min = nums[j]
		} else {
			start = j
		}
	}
	return end - start + 1
}

func Solution2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	l, r := 0, n-1
	for l < n-1 {
		if nums[l] > nums[l+1] {
			break
		}
		l++
	}
	for r > 0 {
		if nums[r] < nums[r-1] {
			break
		}
		r--
	}
	if l >= r {
		return 0
	}
	mi, ma := math.MaxInt32, math.MinInt32
	for i := l; i <= r; i++ {
		mi = min(mi, nums[i])
		ma = max(ma, nums[i])
	}
	for l >= 0 && mi < nums[l] {
		l--
	}
	for r < n && ma > nums[r] {
		r++
	}
	return r - l - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findUnsortedSubarray(nums []int) int {
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(arr)

	i, j := 0, len(arr)-1
	for i < len(nums) {
		if nums[i] == arr[i] {
			i++
		} else {
			break
		}
	}
	for j >= 0 {
		if nums[j] == arr[j] {
			j--
		} else {
			break
		}
	}
	if j < i {
		return 0
	}

	return j - i + 1
}
