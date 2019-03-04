package leetcode

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	maxSum := make([]int, l)
	maxSum[0] = nums[0]
	for i := 1; i < l; i++ {
		if maxSum[i-1]+nums[i] < nums[i] {
			maxSum[i] = nums[i]
		} else {
			maxSum[i] = maxSum[i-1] + nums[i]
		}
	}

	max := maxSum[0]
	for i := 1; i < l; i++ {
		if maxSum[i] > max {
			max = maxSum[i]
		}
	}

	return max
}
