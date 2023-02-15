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
