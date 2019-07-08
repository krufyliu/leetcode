/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */
package leetcode

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	sum := max
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}

	target := max * (max + 1) / 2

	// 缺少最大值或最小值
	if target == sum {
		if len(nums) == max+1 {
			return max + 1
		}
		return 0
	}

	return max*(max+1)/2 - sum
}
