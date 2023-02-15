package Solution

func Solution(s []string) []string {
	start := 0
	end := len(s) - 1

	for start < end {
		s[start], s[end] = s[end], s[start]

		start++
		end--
	}

	return s
}

func reverseString(s []byte) {
	head, end := 0, len(s)-1
	for head < end {
		s[head], s[end] = s[end], s[head]
		head++
		end--
	}
}
