/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */
package leetcode

func majorityElement(nums []int) int {
	counts := map[int]int{}

	n := len(nums)
	for _, v := range nums {
		counts[v]++
		if counts[v] > n/2 {
			return v
		}
	}

	return 0
}
