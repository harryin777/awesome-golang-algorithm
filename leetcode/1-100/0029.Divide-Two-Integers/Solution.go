package Solution

import (
	"math"
)

func divide(dividend int, divisor int) int {
	dm, dsm, sign := true, true, true
	if dividend < 0 {
		dm = false
		dividend = -dividend
	}
	if divisor < 0 {
		dsm = false
		divisor = -divisor
	}

	if (dm && !dsm) || (!dm && dsm) {
		sign = false
	}

	if divisor == 0 {
		return math.MaxInt32
	}

	h := 0
	out := 0
	if dividend < divisor {
		return 0
	}
	for (divisor << uint(h)) <= dividend {
		h++
	}
	h--
	for i := h; i >= 0; i-- {
		if dividend >= (divisor << uint(i)) {
			dividend -= (divisor << uint(i))
			out |= 1 << uint(i)
		}
	}

	if !sign {
		out = -out
	}
	if out > math.MaxInt32 {
		return math.MaxInt32
	} else if out < math.MinInt32 {
		return math.MinInt32
	} else {
		return out
	}
}

// 可行，但是超出时间限制
func divide2(dividend int, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	candidates := []int{divisor}
	for y := divisor; y >= dividend-y; { // 注意溢出
		y += y
		candidates = append(candidates, y)
	}

	ans := 0
	for i := len(candidates) - 1; i >= 0; i-- {
		if candidates[i] >= dividend {
			ans |= 1 << i
			dividend -= candidates[i]
		}
	}
	if rev {
		return -ans
	}
	return ans
}
