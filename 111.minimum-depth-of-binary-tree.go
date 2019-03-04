package leetcode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Right != nil {
		lmh := minDepth(root.Left)
		rmh := minDepth(root.Right)

		if lmh > rmh {
			return 1 + rmh
		}
		return 1 + lmh
	} else if root.Left != nil {
		return 1 + minDepth(root.Left)
	} else if root.Right != nil {
		return 1 + minDepth(root.Right)
	}

	return 1

}
