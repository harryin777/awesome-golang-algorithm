package Solution

func removeElement(nums []int, val int) int {
	tail := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[tail] = nums[i]
			tail++
		}
	}
	return tail
}

func removeElement2(nums []int, val int) int {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

func removeElement3(nums []int, val int) int {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == val {
			right++
		} else {
			if left == right {
				left++
				right++
			} else {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right++
			}
		}
	}

	return left
}
