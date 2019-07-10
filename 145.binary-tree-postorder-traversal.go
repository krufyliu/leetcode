/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcode

func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int

	if root == nil {
		return res
	}

	for {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Right
		}

		if len(stack) == 0 {
			break
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.Left != nil {
			root = top.Left
		}
	}
	// 反转
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return res
}
