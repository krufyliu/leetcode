package leetcode

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	var m int
	for l <= r {
		m = (l + r) >> 1
		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return r + 1
}
