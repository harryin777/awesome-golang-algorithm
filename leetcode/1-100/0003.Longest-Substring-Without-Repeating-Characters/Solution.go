package Solution

import "strings"

func lengthOfLongestSubstring_3(s string) int {
	ans, left, m := 0, 0, map[rune]int{}
	for right, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = right
		} else {
			if m[v]+1 > left {
				left = m[v] + 1
			}
			m[v] = right
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func lengthOfLongestSubstring_1(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func lengthOfLongestSubstring_2(s string) int {
	ans, left, m := 0, 0, map[rune]int{}
	for right, v := range s {
		left = max(left, m[v])
		m[v] = right + 1
		ans = max(ans, right-left+1)
	}
	return ans
}

// O(n) time O(1) space Solution
func lengthOfLongestSubstring33(s string) int {
	var chPosition [256]int // [0, 0, 0, ...]
	maxLength, substringLen, lastRepeatPos := 0, 0, 0

	for i := 0; i < len(s); i++ {
		if pos := chPosition[s[i]]; pos > 0 {
			// record current substring length
			maxLength = Max(substringLen, maxLength)

			// update characters position
			chPosition[s[i]] = i + 1

			// update last repeat character position
			lastRepeatPos = Max(pos, lastRepeatPos)

			// update the current substring from last repeat character
			substringLen = i + 1 - lastRepeatPos
		} else {
			substringLen += 1
			chPosition[s[i]] = i + 1
		}
	}

	return Max(maxLength, substringLen)
}

// 暴力求解(会超时)
func lengthOfLongestSubstring2(s string) int {
	ans := 0

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if allUnique(s, i, j) {
				ans = Max(ans, j-i)
			}
		}
	}
	return ans
}

func allUnique(s string, start int, end int) bool {
	sMap := make(map[string]int)

	for i := start; i < end; i++ {
		if sMap[string(s[i])] > 0 {
			return false
		}
		sMap[string(s[i])]++
	}

	return true
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 1
	ans := ^int(^uint(0) >> 1)
	for right <= len(s) {
		tmp := s[left:right]
		_ = tmp
		c := check(s[left:right])
		if c {
			ans = max(ans, len(s[left:right]))
			right++
		}

		for !c {
			left++
			c = check(s[left:right])
			if c {
				ans = max(ans, len(s[left:right]))
			}
		}
	}

	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func check(s string) bool {
	duplicateMap := make(map[int32]struct{})
	for _, val := range s {
		if _, e := duplicateMap[val]; e {
			return false
		} else {
			duplicateMap[val] = struct{}{}
		}
	}

	return true
}

func lengthOfLongestSubstring3(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	left, right := 0, 1
	var res int
	for right < len(s) {
		tmp := s[left:right]
		if strings.Contains(tmp, string(s[right])) {
			res = max(res, len(tmp))
			left = right
		}
		right++
	}

	res = max(res, len(s[left:right]))

	if left == 0 {
		return len(s)
	}

	return res
}
