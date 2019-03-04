package leetcode

func mySqrt(x int) int {
	l, r := 0, x
	var m int
	for l <= r {
		m = (l + r) >> 1
		mm := m * m
		if mm == x {
			return m
		}

		if mm > x {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return r
}
