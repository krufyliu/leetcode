package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var seq []int
	var stack []*TreeNode

	cur := root

	for {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) == 0 {
			return seq
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		seq = append(seq, top.Val)

		if top.Right != nil {
			cur = top.Right
		}
	}
}
