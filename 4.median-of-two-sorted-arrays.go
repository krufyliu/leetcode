/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (25.51%)
 * Total Accepted:    395.4K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 *
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 *
 * You may assume nums1 and nums2 cannot be both empty.
 *
 * Example 1:
 *
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * The median is 2.0
 *
 *
 * Example 2:
 *
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * The median is (2 + 3)/2 = 2.5
 *
 *
 */
package leetcode

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	after := (m + n - 1) / 2

	lo, hi := 0, m
	for lo < hi {
		i := (lo + hi) / 2
		j := after - i
		cond1 := (j >= 1 && nums1[i] >= nums2[j-1]) || j == 0
		cond2 := (i >= 1 && nums2[j] >= nums1[i-1]) || i == 0
		if cond1 && cond2 {
			lo = i
			break
		} else if !cond1 {
			lo = i + 1
		} else {
			hi = i
		}
	}

	i := lo
	j := after - i
	nextfew := append(nums1[i:i+2], nums2[j:j+2]...)
	sort.Slice(nextfew, func(i, j int) bool {
		return nextfew[i] < nextfew[j]
	})
	return float64(nextfew[0]+nextfew[1-(m+n)%2]) / 2.0
}
