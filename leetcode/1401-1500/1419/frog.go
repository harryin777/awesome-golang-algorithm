package _419

func minNumberOfFrogs(croakOfFrogs string) int {
	if len(croakOfFrogs)%5 != 0 {
		return -1
	}
	m2 := make(map[int]int)
	res := 0
	ans := 0
	for i := 0; i < len(croakOfFrogs); i++ {
		if croakOfFrogs[i] == 'c' {
			m2[i] = 2
			res++
			if res > ans {
				ans = res
			}
			continue
		}
		findNum := 0
		switch croakOfFrogs[i] {
		case 'r':
			findNum = 2
		case 'o':
			findNum = 3
		case 'a':
			findNum = 4
		case 'k':
			findNum = 5
		}

		for key, val := range m2 {
			if val == findNum {
				m2[key]++
				if findNum == 5 {
					res--
				}
				break
			}
		}
	}

	for _, val := range m2 {
		if val != 6 {
			return -1
		}
	}

	return ans
}
