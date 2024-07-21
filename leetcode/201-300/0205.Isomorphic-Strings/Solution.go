package Solution

func Solution(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}

	mapping := make(map[byte]byte)
	reverseMapping := make(map[byte]struct{})
	for idx := 0; idx < ls; idx++ {
		if v, ok := mapping[s[idx]]; ok {
			if v != t[idx] {
				return false
			}
			continue
		}
		if _, ok := reverseMapping[t[idx]]; ok {
			return false
		}

		mapping[s[idx]] = t[idx]
		reverseMapping[t[idx]] = struct{}{}
	}

	return true
}

func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if (m1[s[i]] > 0 && m1[s[i]] != t[i]) || (m2[t[i]] > 0 && m2[t[i]] != s[i]) {
			return false
		}

		m1[s[i]] = t[i]
		m2[t[i]] = s[i]
	}

	return true
}
