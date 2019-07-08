/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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

func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int

	if root == nil {
		return res
	}

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) == 0 {
			break
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		if top.Right != nil {
			root = top.Right
		}
	}

	return res
}
