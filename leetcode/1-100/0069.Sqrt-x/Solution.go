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
	l, r := 1, x
	for r-l > 1 {
		mid := l + (r-l)/2
		product := mid * mid
		if product == x {
			return mid
		} else if product < x {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}
