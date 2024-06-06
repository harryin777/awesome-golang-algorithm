package Solution

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	s1cnt, index, s2cnt := 0, 0, 0
	recall := make(map[int][2]int)
	var preLoop, inLoop [2]int

	for {
		s1cnt++
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[index] {
				index++
				if index == len(s2) {
					s2cnt++
					index = 0
				}
			}
		}

		if s1cnt == n1 {
			return s2cnt / n2
		}

		if value, ok := recall[index]; ok {
			s1cntPrime := value[0]
			s2cntPrime := value[1]
			preLoop = [2]int{s1cntPrime, s2cntPrime}
			inLoop = [2]int{s1cnt - s1cntPrime, s2cnt - s2cntPrime}
			break
		} else {
			recall[index] = [2]int{s1cnt, s2cnt}
		}
	}

	ans := preLoop[1] + (n1-preLoop[0])/inLoop[0]*inLoop[1]
	rest := (n1 - preLoop[0]) % inLoop[0]

	for i := 0; i < rest; i++ {
		for j := 0; j < len(s1); j++ {
			if s1[j] == s2[index] {
				index++
				if index == len(s2) {
					ans++
					index = 0
				}
			}
		}
	}

	return ans / n2
}
