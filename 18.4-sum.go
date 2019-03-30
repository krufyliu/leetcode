/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (29.64%)
 * Total Accepted:    216.9K
 * Total Submissions: 725.7K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers and an integer target, are there elements
 * a, b, c, and d in nums such that a + b + c + d = target? Find all unique
 * quadruplets in the array which gives the sum of target.
 *
 * Note:
 *
 * The solution set must not contain duplicate quadruplets.
 *
 * Example:
 *
 *
 * Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
 *
 * A solution set is:
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */
package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	// l := len(nums)
	// var solutions [][]int

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return findNsum(nums, target, 4, nil)

	// var i, j, k, e int
	// for i < l-3 {
	// 	e = l - 1
	// 	for e-i >= 3 {
	// 		j = i + 1
	// 		k = e - 1
	// 		psum := nums[i] + nums[e]
	// 		for j < k {
	// 			sum := psum + nums[j] + nums[k]
	// 			if sum < target {
	// 				j++
	// 				continue
	// 			} else if sum > target {
	// 				k--
	// 				continue
	// 			}
	// 			solutions = append(solutions, []int{nums[i], nums[j], nums[k], nums[e]})
	// 			j++
	// 			k--
	// 			for j < k && nums[j-1] == nums[j] {
	// 				j++
	// 			}
	// 			for j < k && nums[k+1] == nums[k] {
	// 				k--
	// 			}
	// 		}
	// 		e--
	// 		for e-i >= 3 && nums[e+1] == nums[e] {
	// 			e--
	// 		}
	// 	}
	// 	i++
	// 	for i < l-3 && nums[i-1] == nums[i] {
	// 		i++
	// 	}
	// }

	// return solutions
}

// nums is a sorted array
func findNsum(nums []int, target int, n int, res []int) [][]int {
	l := len(nums)
	if l < n || n < 2 {
		return nil
	}

	if n == 2 {
		return findTwoSum(nums, target)
	}

	var solutions [][]int
	for i := 0; i <= l-n; i++ {
		if target < nums[i]*n || nums[l-1]*n < target {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			parts := findNsum(nums[i+1:], target-nums[i], n-1, append(res, nums[i]))
			if n == 3 {
				for j := 0; j < len(parts); j++ {
					nres := make([]int, len(res), n)
					copy(nres, res)
					nres = append(nres, nums[i])
					nres = append(nres, parts[j]...)
					solutions = append(solutions, nres)
				}
			} else {
				solutions = append(solutions, parts...)
			}
		}
	}

	return solutions
}

func findTwoSum(nums []int, target int) [][]int {
	l, r := 0, len(nums)-1
	var solutions [][]int
	for l < r {
		sum := nums[l] + nums[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			solutions = append(solutions, []int{nums[l], nums[r]})
			l++
			r--
			for l < r && nums[l-1] == nums[l] {
				l++
			}
			for l < r && nums[r+1] == nums[r] {
				r--
			}
		}
	}

	return solutions
}
