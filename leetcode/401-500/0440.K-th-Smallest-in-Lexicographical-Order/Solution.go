package Solution

import (
	"sort"
	"strconv"
)

func findKthNumber(n int, k int) int {
	arr := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		arr = append(arr, strconv.Itoa(i))
	}
	sort.Slice(arr, func(i, j int) bool {
		index := 0
		for index < len(arr[i]) && index < len(arr[j]) {
			if arr[i][index] < arr[j][index] {
				return true
			} else if arr[i][index] > arr[j][index] {
				return false
			} else {
				index++
			}
		}
		if index >= len(arr[i]) {
			return true
		} else {
			return false
		}
	})
	ans, _ := strconv.Atoi(arr[k-1])
	return ans

}
