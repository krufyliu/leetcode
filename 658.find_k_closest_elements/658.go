package find_k_closest_elements

func FindClosestElements(arr []int, k int, x int) []int {
	length := len(arr)
	if x <= arr[0] {
		return arr[0:k]
	}
	if x >= arr[length-1] {
		return arr[length-k:]
	}

	var l, r = 0, length - 1
	var m int
	for l <= r {
		m = (l + r) / 2
		if arr[m] >= x {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	var t = make([]int, k)
	var count, i, j = 0, m, m + 1
	for count < k {
		count++
		if i < 0 {
			t[k-count] = arr[j]
			j++
			continue
		}

		if j >= length {
			t[k-count] = arr[i]
			i--
			continue
		}

		if x-arr[i] <= arr[j]-x {
			t[k-count] = arr[i]
			i--
		} else {
			t[k-count] = arr[j]
			j++
		}

	}

	return t
}
