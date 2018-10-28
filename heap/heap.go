package heap

func heapfiy(arr []int) {
	l := len(arr)

	if l < 2 {
		return
	}

	for i := l/2 - 1; i >= 0; i-- {
		down(arr, i, l)
	}
}

func up(arr []int, i int) {
	for {
		j := (i - 1) / 2
		if i == j || arr[i] >= arr[j] {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i = j
	}

}

func down(arr []int, i0, n int) {
	var lc, rc, si int

	i := i0

	for {
		lc = 2*i + 1
		if lc >= n {
			break
		}
		rc = lc + 1
		si = i
		if arr[i] > arr[lc] {
			si = lc
		}
		if rc < n && arr[si] > arr[rc] {
			si = rc
		}
		if si == i {
			break
		}
		arr[si], arr[i] = arr[i], arr[si]
		i = si
	}
}
