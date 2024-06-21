package Solution

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] > nums[mid-1] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func findPeakElement2(nums []int) int {
	len := len(nums)
	left, right := 0, len-1
	for left < right {
		mid := left + (right-left)/2
		if mid > 0 && mid < len-1 {
			if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
				return mid
			}
			if nums[mid] > nums[mid-1] {
				left++
			} else {
				right--
			}
		}
	}

	return 0
}
