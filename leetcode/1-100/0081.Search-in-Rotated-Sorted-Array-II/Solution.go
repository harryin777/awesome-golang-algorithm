package Solution

func search(nums []int, target int) bool {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return true
		}

		if nums[mid] == nums[left] {
			left++
			continue
		}

		if nums[mid] > nums[left] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	if left >= len(nums) {
		return false
	}

	return nums[left] == target
}
