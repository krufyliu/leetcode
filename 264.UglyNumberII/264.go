package UglyNumberII

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	var i, next int

	var uglyNums = make([]int, n)
	var p2, p3, p5 int
	uglyNums[0] = 1

	for i < n-1 {
		next = min(uglyNums[p2]*2, min(uglyNums[p3]*3, uglyNums[p5]*5))
		i++
		uglyNums[i] = next
		if uglyNums[p2]*2 == next {
			p2++
		}
		if uglyNums[p3]*3 == next {
			p3++
		}
		if uglyNums[p5]*5 == next {
			p5++
		}
	}
	return uglyNums[n-1]
}
