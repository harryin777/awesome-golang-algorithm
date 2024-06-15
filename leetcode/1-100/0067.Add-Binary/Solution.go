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
	pre := 0
	x, y := len(a)-1, len(b)-1
	var res string
	for x >= 0 && y >= 0 {
		sub := pre + int(a[x]-'0'+b[y]-'0')
		if sub > 1 {
			pre = 1
			res = strconv.Itoa(sub%2) + res

		} else {
			res = strconv.Itoa(sub) + res
			pre = 0
		}
		x--
		y--
	}
	if x < 0 && y >= 0 {
		if pre != 0 {
			for y >= 0 {
				sub := pre + int(b[y]-'0')
				if sub > 1 {
					pre = 1
					res = strconv.Itoa(sub%2) + res
				} else {
					res = strconv.Itoa(sub) + res
					pre = 0
				}
				y--
			}
		} else {
			if b[:y+1] == "0" {
				return res
			}
			res = b[:y+1] + res
		}

	} else if y < 0 && x >= 0 {
		if pre != 0 {
			for x >= 0 {
				sub := pre + int(a[x]-'0')
				if sub > 1 {
					pre = 1
					res = strconv.Itoa(sub%2) + res
				} else {
					res = strconv.Itoa(sub) + res
					pre = 0
				}
				x--
			}
		} else {
			if a[:x+1] == "0" {
				return res
			}
			res = a[:x+1] + res
		}

	}
	if pre != 0 {
		res = strconv.Itoa(pre) + res
	}

	return res
}
