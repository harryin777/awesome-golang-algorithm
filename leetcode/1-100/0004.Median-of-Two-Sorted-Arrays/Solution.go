package Solution

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	halfLen := (m + n + 1) / 2

	if m == 0 || n == 0 {
		nums := append(nums1, nums2...)
		if len(nums)%2 == 1 {
			return float64(nums[halfLen-1])
		} else {
			return (float64(nums[halfLen-1]) + float64(nums[halfLen])) / 2.0
		}
	}

	// make sure m < n, so j = halfLen - i is always greater than zero
	var A, B []int
	if m < n {
		A, B = nums1, nums2
	} else {
		m, n = n, m
		A, B = nums2, nums1
	}

	// find i in [0, m]
	// especial, if i == 0 means A is the right part, if i == m means A is the left part
	iMin, iMax := 0, m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i

		if i > 0 && j >= 0 && j < n && A[i-1] > B[j] {
			// Ai > Bj+1, i need be smaller
			iMax = i - 1
		} else if j > 0 && i >= 0 && i < m && B[j-1] > A[i] {
			// Bj > Ai+1, i need be greater
			iMin = i + 1
		} else {
			var leftPartMax, rightPartMin float64
			if i == 0 {
				leftPartMax = float64(B[j-1])
			} else if j == 0 {
				leftPartMax = float64(A[i-1])
			} else {
				leftPartMax = float64(Max(A[i-1], B[j-1]))
			}

			if (m+n)%2 == 1 {
				// m + n is odd
				return leftPartMax
			}

			if i == m {
				rightPartMin = float64(B[j])
			} else if j == n {
				rightPartMin = float64(A[i])
			} else {
				rightPartMin = float64(Min(A[i], B[j]))
			}

			// m + n is even
			return (leftPartMax + rightPartMin) / 2.0
		}
	}

	return -1.0
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m == 0 && n == 0 {
		return 0
	}
	nums := make([]int, 0, m+n)
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)
	nums3 := quickSort(nums, 0, m+n-1)
	if (m+n)%2 == 1 {
		return float64(nums3[(m+n)/2])
	} else {

		return (float64(nums3[(m+n)/2]) + float64(nums3[(m+n)/2-1])) / 2
	}
}

func quickSort(arr []int, left, right int) []int {
	if left > right {
		return arr
	}
	i, j := left, right
	min := arr[i]
	for i < j {
		for i < j && arr[j] >= min {
			j--
		}
		for i < j && arr[i] <= min {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left] = arr[i]
	arr[i] = min

	arr = quickSort(arr, left, i-1)
	arr = quickSort(arr, i+1, right)

	return arr
}

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	arr := merge2(nums1, nums2)
	if len(arr)%2 == 0 {
		mid := len(arr) / 2
		return float64(arr[mid]+arr[mid-1]) / 2
	} else {
		mid := len(arr) / 2
		return float64(arr[mid])
	}

}

func merge2(left, right []int) []int {
	i, j := 0, 0
	res := make([]int, 0, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}
