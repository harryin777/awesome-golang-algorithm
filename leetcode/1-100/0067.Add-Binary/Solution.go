package Solution

import (
	"strconv"
	"strings"
)

func addBinary(a, b string) string {
	lenA := len(a)
	lenB := len(b)
	if lenB > lenA {
		a, b = b, a             // swap
		lenA, lenB = lenB, lenA // swap2
	}
	b = strings.Repeat("0", lenA-lenB) + b
	lenB = len(b)

	carry := byte(0)
	ret := make([]byte, lenB+1)

	for i := lenB - 1; i >= 0; i-- {
		numA := a[i] - '0'
		numB := b[i] - '0'
		sum := numA + numB + carry
		ret[i+1] = sum&1 + '0'
		carry = sum >> 1
	}
	if carry == 0 {
		ret = ret[1:]
	} else {
		ret[0] = carry + '0'
	}
	return string(ret)
}

func addBinary2(a string, b string) string {
	var res string
	m, n := len(a)-1, len(b)-1
	add := 0
	for ; m >= 0 || n >= 0 || add > 0; m, n = m-1, n-1 {
		var x, y int
		if m >= 0 {
			x = int(a[m] - '0')
		}
		if n >= 0 {
			y = int(b[n] - '0')
		}
		sum := add + x + y
		if sum == 3 {
			res = "1" + res
			add = 1
		} else if sum == 2 {
			res = "0" + res
			add = 1
		} else if sum == 1 {
			res = "1" + res
			add = 0
		} else {
			res = "0" + res
			add = 0
		}
	}

	return res
}

func addBinary3(a string, b string) string {
	m, n := len(a)-1, len(b)-1
	add := 0
	ans := ""

	for ; m >= 0 || n >= 0 || add != 0; m, n = m-1, n-1 {
		var x, y int
		if m >= 0 {
			x = int(a[m] - '0')
		}
		if n >= 0 {
			y = int(b[n] - '0')
		}
		sum := 0
		if x+y+add >= 2 {
			sum = (x + y + add) % 2
			add = 1
		} else {
			sum = x + y + add
			add = 0
		}

		ans = strconv.Itoa(sum) + ans
	}

	return ans
}
