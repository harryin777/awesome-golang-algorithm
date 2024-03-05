package Solution

func openLock(deadends []string, target string) int {
	deadendMap := make(map[string]struct{})
	for _, deadend := range deadends {
		deadendMap[deadend] = struct{}{}
	}

	visited := make(map[string]struct{})
	q := make([]string, 0, 10)
	q = append(q, "0000")
	step := 0
	visited["0000"] = struct{}{}
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if cur == target {
				return step
			}

			if _, ok := deadendMap[cur]; ok {
				continue
			}
			for j := 0; j < 4; j++ {
				up := plusOne(cur, j)
				if _, ok := visited[up]; !ok {
					q = append(q, up)
					visited[up] = struct{}{}
				}
				down := minusOne(cur, j)
				if _, ok := visited[down]; !ok {
					q = append(q, down)
					visited[down] = struct{}{}
				}
			}

		}
		step++
	}

	return -1
}

func plusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j] += 1
	}
	return string(ch)
}

// 将 s[i] 向下拨动一次
func minusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j] -= 1
	}
	return string(ch)
}
