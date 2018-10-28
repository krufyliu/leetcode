package MaximumSumCircularSubarray

func maxSubarraySumCircular(A []int) int {
	var l = len(A)
	var max = A[0]
	var maxAll = A[0]
	var startI = 0
	var i = 1

	for i < 2*l-1 {
		if i-startI+1 > l {
			max = A[i%l]
			startI = i
		} else if max < 0 {
			max = A[i%l]
			startI = i
		} else {
			max += A[i%l]
		}
		if max > maxAll {
			maxAll = max
		}
		i++
	}

	return maxAll
}
