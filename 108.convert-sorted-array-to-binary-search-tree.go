package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)

	if l == 0 {
		return nil
	}

	if l == 1 {
		return &TreeNode{Val: nums[0]}
	}

	m := l / 2

	root := &TreeNode{Val: nums[m]}
	root.Left = sortedArrayToBST(nums[0:m])
	root.Right = sortedArrayToBST(nums[m+1:])

	return root
}
