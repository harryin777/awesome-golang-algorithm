package Solution

func combinationSum3(k int, n int) (result [][]int) {
	dfs(&result, make([]int, 0), n, 1, k)
	return
}

func dfs(result *[][]int, comb []int, remain int, num int, length int) {
	if remain < 0 || len(comb) > length {
		return
	}
	if remain == 0 && len(comb) == length {
		combCopy := make([]int, len(comb))
		copy(combCopy, comb)
		*result = append(*result, combCopy)
		return
	}

	for i := num; i <= 9; i++ {
		comb = append(comb, i)
		dfs(result, comb, remain-i, i+1, length)
		comb = comb[:len(comb)-1]
	}
}

func combinationSum4(k int, n int) (ans [][]int) {
	var temp []int
	check := func(mask int) bool {
		temp = nil
		sum := 0
		for i := 0; i < 9; i++ {
			if 1<<i&mask > 0 {
				temp = append(temp, i+1)
				sum += i + 1
			}
		}
		return len(temp) == k && sum == n
	}

	for mask := 0; mask < 1<<9; mask++ {
		if check(mask) {
			ans = append(ans, append([]int(nil), temp...))
		}
	}
	return
}
