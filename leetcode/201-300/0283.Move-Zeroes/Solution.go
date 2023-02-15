package Solution

func moveZeroes(nums []int) []int {
	c := 0
	for _, v := range nums {
		if v != 0 {
			nums[c] = v
			c++
		}
	}
	for ; c < len(nums); c++ {
		nums[c] = 0
	}
	return nums
}

func moveZeroes_2(nums []int) []int {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	return nums
}

func moveZeroes_3(nums []int) []int {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}

	return nums
}
