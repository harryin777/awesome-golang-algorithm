package Solution

func removeDuplicates(s string, k int) string {
	left, right := 0, 1
	arr := []byte(s)
	for right < len(arr) {
		if arr[left] != arr[right] {
			left++
			right++
			continue
		}
		if right-left+1 == k {
			arr = append(arr[:left], arr[right+1:]...)
			right = left
			left--
		}
	}

	return string(arr)
}
