package Solution

func Solution(nums []int) int {
	cmpObj, repeatCount := nums[0], 1
	fillIdx := 0
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] == cmpObj {
			repeatCount++
			continue
		}

		cmpObj = nums[idx]

		if repeatCount > 2 {
			repeatCount = 2
		}
		preObj := nums[idx-1]
		for ; repeatCount > 0; repeatCount, fillIdx = repeatCount-1, fillIdx+1 {
			nums[fillIdx] = preObj
		}
		repeatCount = 1
	}
	if repeatCount > 2 {
		repeatCount = 2
	}

	preObj := nums[len(nums)-1]
	for ; repeatCount > 0; repeatCount, fillIdx = repeatCount-1, fillIdx+1 {
		nums[fillIdx] = preObj
	}

	return fillIdx
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	left, right := 0, 1
	for right < len(nums) {
		if nums[left] == nums[right] {
			if right-left != 1 {
				nums = append(nums[:left+1], nums[right:]...)
				right = left + 1
			}
			right++

		} else {
			left = right
			right++
		}
	}

	return len(nums)
}
