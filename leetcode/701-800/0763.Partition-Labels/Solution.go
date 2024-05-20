package Solution

func Solution(s string) []int {
	partLength := make([]int, 0)
	charsLastIndex := make([]int, 26)
	for idx, b := range s {
		charsLastIndex[b-'a'] = idx
	}

	lastIdx, length := -1, 0
	for idx := 0; idx < len(s); idx++ {
		if charsLastIndex[s[idx]-'a'] > lastIdx {
			lastIdx = charsLastIndex[s[idx]-'a']
		}
		if idx == lastIdx {
			partLength = append(partLength, length+1)
			length = 0
			continue
		}
		length++
	}
	return partLength
}

func partitionLabels(s string) []int {
	lastPos := make([]int, 26)
	for i := 0; i < len(s); i++ {
		lastPos[s[i]-'a'] = i
	}

	res := make([]int, 0, 10)
	end := 0
	count := 0
	for i := 0; i < len(s); i++ {
		end = max(end, lastPos[s[i]-'a'])
		if i == end {
			res = append(res, count+1)
			count = 0
			continue
		}
		count++
	}

	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
