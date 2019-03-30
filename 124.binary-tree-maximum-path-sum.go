package leetcode

func maxPathSum(root *TreeNode) int {
	_, max := pathSumWithMax(root)
	return max
}

func pathSumWithMax(root *TreeNode) (int, int) {
	sum, maxSum := root.Val, root.Val

	if root.Left == nil && root.Right == nil {
		return sum, maxSum
	}

	if root.Left != nil && root.Right != nil {
		lsum, lmax := pathSumWithMax(root.Left)
		rsum, rmax := pathSumWithMax(root.Right)
		if root.Val < 0 {
			if maxSum < lmax {
				maxSum = lmax
			}
			if maxSum < rmax {
				maxSum = rmax
			}

			if maxSum < root.Val+lsum+rsum {
				maxSum = root.Val + lsum + rsum
			}
		} else {
			if lsum > 0 {
				maxSum += max(lsum, root.Left.Val)
				if maxSum < lmax {
					maxSum = lmax
				}
			}
			if rsum > 0 {
				maxSum += max(rsum, root.Right.Val)
				if maxSum < rmax {
					maxSum = rmax
				}
			}
		}
		sum += lsum
		sum += rsum

	} else if root.Left != nil {
		lsum, lmax := pathSumWithMax(root.Left)
		if root.Val < 0 {
			if maxSum < lmax {
				maxSum = lmax
			}
		} else {
			if lsum > 0 {
				maxSum += max(lsum, root.Left.Val)
				if maxSum < lmax {
					maxSum = lmax
				}
			}
		}
		sum += lsum
	} else {
		rsum, rmax := pathSumWithMax(root.Right)
		if root.Val < 0 {
			if maxSum < rmax {
				maxSum = rmax
			}
		} else {
			if rsum > 0 {
				maxSum += max(rsum, root.Right.Val)
				if maxSum < rmax {
					maxSum = rmax
				}
			}
		}
		sum += rsum
	}

	return sum, maxSum
}

func maxPathSumAndHalf(root *TreeNode) (int, int) {
	return 0, 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
