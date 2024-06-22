package Solution

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")
	l := max(len(arr1), len(arr2))
	for i := 0; i < l; i++ {
		if i < len(arr1) && i < len(arr2) {
			v1, _ := strconv.ParseInt(arr1[i], 10, 64)
			v2, _ := strconv.ParseInt(arr2[i], 10, 64)
			if v1 > v2 {
				return 1
			} else if v1 < v2 {
				return -1
			}
		} else if i >= len(arr1) {
			var v1 int64
			v2, _ := strconv.ParseInt(arr2[i], 10, 64)
			if v1 > v2 {
				return 1
			} else if v1 < v2 {
				return -1
			}
		} else {
			var v2 int64
			v1, _ := strconv.ParseInt(arr1[i], 10, 64)
			if v1 > v2 {
				return 1
			} else if v1 < v2 {
				return -1
			}
		}
	}

	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
