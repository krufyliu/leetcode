/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */
package leetcode

func twoSum(numbers []int, target int) []int {
	l := len(numbers)

	l, r := 1, l
	for l < r {
		sum := numbers[l-1] + numbers[r-1]
		if sum == target {
			return []int{l, r}
		}

		if sum > target {
			r = r - 1
		} else {
			l = l + 1
		}
	}

	return nil
}
