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
		if i > 1 && fruits[i] == fruits[i-1] {
			continue
		}
		m := make(map[int]struct{})
		var j = i
		for ; j < len(fruits); j++ {
			m[fruits[j]] = struct{}{}
			if len(m) > 2 {
				break
			}
		}
		ans = max(ans, j-i)
	}

	return ans
}
