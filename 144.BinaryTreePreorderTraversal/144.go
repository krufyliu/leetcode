package BinaryTreePreorderTraversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack []*TreeNode
	var seq []int

	for {
		for root != nil {
			seq = append(seq, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) == 0 {
			return seq
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.Right != nil {
			root = top.Right
		}
	}
}
