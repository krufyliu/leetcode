/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
package leetcode

func search(nums []int, target int) int {
	ll := len(nums)
	l, r := 0, ll-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		// nums[m] > nums[0]
		// nums[m] < nums[ll-1]
		// normal r = m - 1
		//
		if nums[m] > target {
			if nums[m] <= nums[r] || nums[l] <= target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[l] <= nums[m] || target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return -1
}
