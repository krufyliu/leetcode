/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */
package leetcode

func findKthLargest(nums []int, k int) int {
	p := partion(nums)

	if p == k-1 {
		return nums[p]
	}

	if p < k-1 {
		return findKthLargest(nums[p+1:], k-p-1)
	}

	return findKthLargest(nums[0:p], k)
}

func partion(nums []int) int {
	if len(nums) < 2 {
		return len(nums) - 1
	}

	guard := nums[0]

	i, j := 1, len(nums)-1
	for i <= j {
		for i <= j && nums[i] > guard {
			i++
		}

		for i <= j && nums[j] < guard {
			j--
		}

		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	nums[0], nums[j] = nums[j], nums[0]

	return j
}
