/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	solutions := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ll := len(nums)

	for i := 0; i < ll-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			parts := twoSum15(nums[i+1:], -nums[i])
			for j := 0; j < len(parts); j++ {
				solutions = append(solutions, []int{nums[i], parts[j][0], parts[j][1]})
			}
		}
	}

	return solutions
}

func twoSum15(nums []int, target int) [][]int {
	solutions := [][]int{}
	ll := len(nums)
	i, j := 0, ll-1
	for i < j {
		if nums[i]+nums[j] > target {
			j--
			continue
		} else if nums[i]+nums[j] < target {
			i++
			continue
		}
		solutions = append(solutions, []int{nums[i], nums[j]})
		i++
		j--
		for i < j && nums[i] == nums[i-1] {
			i++
		}
		for i < j && nums[j] == nums[j+1] {
			j--
		}
	}

	return solutions
}
