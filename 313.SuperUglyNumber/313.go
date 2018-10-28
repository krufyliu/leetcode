package SuperUglyNumber

func min(arr []int) int {
	mi := 0
	mv := arr[0]
	for i, v := range arr[0:] {
		if v < mv {
			mv = v
			mi = i
		}
	}
	return mi
}

func next(uglyNums []int, pos []int, primes []int) []int {
	ns := make([]int, len(pos))
	for i, v := range pos {
		ns[i] = uglyNums[v] * primes[i]
	}
	return ns
}

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}

	uglyNums := make([]int, n)

	pos := make([]int, len(primes))

	uglyNums[0] = 1

	var i int
	for i < n-1 {
		ns := next(uglyNums, pos, primes)
		mi := min(ns)
		pos[mi]++
		if ns[mi] == uglyNums[i] {
			continue
		}
		i++
		uglyNums[i] = ns[mi]
	}
	return uglyNums[n-1]
}
