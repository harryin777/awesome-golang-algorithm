package Solution

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	tail := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[tail] = nums[i]
			tail++
		}
	}
	return tail
}

func removeDuplicates2(nums []int) int {
	fast, slow := 1, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	fmt.Println(nums)
	return slow + 1
}

func removeDuplicates3(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	slow, quick := 0, 1
	for quick < len(nums) {
		if nums[slow] == nums[quick] {
			if quick == len(nums)-1 {
				nums = nums[:quick]
			} else {
				nums = append(nums[:slow+1], nums[slow+2:]...)
			}
		} else {
			slow++
			quick++
		}
	}

	return len(nums)
}

func removeDuplicates4(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	slow, quick := 0, 1
	for quick < len(nums) {
		if nums[slow] != nums[quick] {
			if quick-slow > 1 {
				nums[slow+1] = nums[quick]
			}
			slow++
			quick++
		} else {
			quick++
		}
	}

	return slow
}
