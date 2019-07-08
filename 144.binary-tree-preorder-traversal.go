/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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

func preorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int

	if root == nil {
		return res
	}

	for {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) == 0 {
			break
		}
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if top.Right != nil {
			root = top.Right
		}
	}

	return res
}
