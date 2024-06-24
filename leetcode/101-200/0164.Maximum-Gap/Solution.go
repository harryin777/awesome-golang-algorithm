package Solution

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	arr := merge1(nums)
	i, j, res := 0, 1, 0
	for j <= len(arr)-1 {
		if arr[j]-arr[i] > res {
			res = arr[j] - arr[i]
		}
		i++
		j++
	}
	if i == 0 {
		return arr[j] - arr[i]
	}

	return res
}

func merge1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := merge1(arr[:mid])
	right := merge1(arr[mid:])

	return merge2(left, right)
}

func merge2(left, right []int) []int {
	i, j := 0, 0
	res := make([]int, 0, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
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
