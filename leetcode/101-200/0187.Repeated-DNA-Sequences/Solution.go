package Solution

func findRepeatedDnaSequences(s string) []string {
	res := make([]string, 0, 10)
	dupMap := make(map[string]int)
	dupMap2 := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		if _, ok := dupMap[s[i:i+10]]; ok {
			dupMap2[s[i:i+10]] = 1
		} else {
			dupMap[s[i:i+10]] = 1
		}
	}

	for key, _ := range dupMap2 {
		res = append(res, key)
	}

	return res
}
