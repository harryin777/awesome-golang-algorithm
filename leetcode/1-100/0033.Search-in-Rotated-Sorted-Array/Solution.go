package Solution

func search(nums []int, target int) int {
	low, mid, high := 0, 0, len(nums)-1
	for low <= high {
		mid = (low + high) / 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[low] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[l] {
			if nums[l] <= target && target < nums[mid] {
				r--
			} else {
				l++
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l++
			} else {
				r--
			}
		}
	}

	return -1
}
