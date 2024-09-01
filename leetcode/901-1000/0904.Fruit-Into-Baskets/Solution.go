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
	count := make(map[int]int)
	left := 0
	for index, x := range fruits {
		count[x]++
		for len(count) > 2 {
			l := fruits[left]
			count[l]--
			if count[l] == 0 {
				delete(count, l)
			}
			left++
		}
		ans = max(ans, index-left+1)
	}

	return ans
}
