/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
package leetcode

func searchRange(nums []int, target int) []int {
	p := searchTarget(nums, target)
	if p == -1 {
		return []int{-1, -1}
	}

	pl, pr := p-1, p+1
	// 往走遍找
	for pl >= 0 && nums[pl] == target {
		pl--
	}

	// 往右边找
	for pr < len(nums) && nums[pr] == target {
		pr++
	}

	return []int{pl + 1, pr - 1}
}

func searchTarget(nums []int, target int) int {
	n := len(nums)

	l, r := 0, n-1
	for l <= r {
		m := (l + r) >> 1
		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}
