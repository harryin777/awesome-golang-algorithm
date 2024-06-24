package Solution

func Solution(nums []int, k int) {
	n := len(nums)
	if k == 0 || k%n == 0 {
		return
	}
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(arr []int, x, y int) {
	for i, j := x, y; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func rotate(nums []int, k int) {
	l := len(nums)
	reminder := k % l
	if reminder == 0 {
		return
	}
	tmp2 := nums
	tmp := make([]int, reminder)
	copy(tmp, tmp2[l-reminder:])
	tmp2 = tmp2[:l-reminder]
	tmp2 = append(tmp, tmp2...)
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp2[i]
	}

	return
}
