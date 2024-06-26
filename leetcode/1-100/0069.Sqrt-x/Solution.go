package Solution

func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

func mySqrt2(x int) int {
	if x == 1 || x == 0 {
		return x
	}

	l, r, precision := 0, x, 1
	for (r - l) > precision {
		mid := l + (r-l)/2
		power := mid * mid
		if power > x {
			r = mid
		} else {
			l = mid
		}
	}

	return l
}
