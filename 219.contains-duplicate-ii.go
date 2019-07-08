/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */
package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	v2indexes := map[int][]int{}

	for i, v := range nums {
		indexes := v2indexes[v]
		if len(indexes) != 0 {
			for _, j := range indexes {
				diff := i - j
				if diff <= k {
					return true
				}
			}

		}
		v2indexes[v] = append(indexes, i)
	}

	return false
}
