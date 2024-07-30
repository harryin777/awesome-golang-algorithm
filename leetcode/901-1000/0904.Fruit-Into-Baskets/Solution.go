package Solution

func totalFruit_1(fruits []int) int {
	ans, left, win := 0, 0, map[int]int{}
	for right, v := range fruits {
		win[v]++
		for len(win) > 2 {
			win[fruits[left]]--
			if win[fruits[left]] == 0 {
				delete(win, fruits[left])
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func totalFruit(fruits []int) int {
	ans := 0
	for i := 0; i < len(fruits); i++ {
		m := make(map[int]struct{})
		for j := 0; j < len(fruits); j++ {
			m[fruits[j]] = struct{}{}
			if len(m) > 2 {
				ans = max(ans, j-1)
				break
			}
		}
	}

	return ans
}
