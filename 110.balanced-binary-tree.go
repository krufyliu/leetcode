package leetcode

func isBalanced(root *TreeNode) bool {
	h := depthV2(root)
	return h != -1
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := depth(root.Left)
	rh := depth(root.Right)

	if lh > rh {
		return 1 + lh
	}

	return 1 + rh
}

func depthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := depthV2(root.Left)
	if lh == -1 {
		return -1
	}
	rh := depthV2(root.Right)
	if rh == -1 {
		return -1
	}

	if lh-rh > 1 || rh-lh > 1 {
		return -1
	}

	if lh > rh {
		return 1 + lh
	}

	return 1 + rh
}
