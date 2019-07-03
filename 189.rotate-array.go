/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */
package leetcode

func rotate(nums []int, k int) {
	rotateByCyclic(nums, k)
}

func rotateUsingExtraSpace(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	k = k % n
	if k == 0 {
		return
	}

	newnums := make([]int, n)
	copy(newnums[:k], nums[n-k:])
	copy(newnums[k:], nums[:n-k])
	copy(nums[0:], newnums[0:])
}

func rotateByCyclic(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	k = k % n
	if k == 0 {
		return
	}

	count := 0
	for i := 0; count < n; i++ {
		cur := i
		next := (i + k) % n
		tmp := nums[cur]
		for next != i {
			nums[next], tmp = tmp, nums[next]
			next = (next + k) % n
			count++
		}
		nums[next] = tmp
		count++
	}
}
