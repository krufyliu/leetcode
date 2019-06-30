/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */
package leetcode

func nextPermutation(nums []int) {
	l := len(nums)
	i := l - 2
	// 从右向左找到非递增的数字
	for i >= 0 {
		if nums[i+1] > nums[i] {
			break
		}
		i--
	}

	if i >= 0 {
		// 从右向左遍历，找到第一个比nums[i] 大的数
		j := l - 1
		for nums[j] <= nums[i] {
			j--
		}
		// 交换两个数，nums[i]右边还是降序的
		// 即将i位置填充下一个比nums[i]大的数
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 将nums[i]右边的序列从降序变成升序即使下一个最下的permutation
	j, k := i+1, l-1
	for j < k {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
}
