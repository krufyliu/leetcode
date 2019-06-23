/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
package leetcode

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ll := len(nums)

	targetSum := nums[0] + nums[1] + nums[2]
	closest := target - targetSum
	if closest < 0 {
		closest = -closest
	}

	for i := 0; i < ll-2; i++ {
		j := i + 1
		k := ll - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			// 差值向0靠近
			diff := target - sum
			if diff > 0 {
				j++
				if diff < closest {
					closest = diff
					targetSum = sum

				}
			} else {
				k--
				if -diff < closest {
					closest = -diff
					targetSum = sum
				}
			}
		}
	}

	return targetSum
}
