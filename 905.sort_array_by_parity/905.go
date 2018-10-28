package sort_array_by_parity

func sortArrayByParity(A []int) []int {
	var i, j int = 0, len(A) - 1
	for i < j {
		for A[i]&1 == 0 {
			i++
		}
		for A[j]&1 == 1 {
			j--
		}
		if i < j {
			A[i], A[j] = A[j], A[i]
			i++
			j--
		}
	}
	return A
}
