/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (37.87%)
 * Total Accepted:    145K
 * Total Submissions: 377.5K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * Given inorder and postorder traversal of a tree, construct the binary tree.
 *
 * Note:
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * inorder = [9,3,15,20,7]
 * postorder = [9,15,7,20,3]
 *
 * Return the following binary tree:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	l := len(postorder)
	if l == 0 {
		return nil
	}

	first := postorder[l-1]
	root := &TreeNode{
		Val: first,
	}

	if l == 1 {
		return root
	}

	p := 0
	for p < l {
		if inorder[p] == first {
			break
		}
		p++
	}
	root.Left = buildTree(inorder[0:p], postorder[:p])
	root.Right = buildTree(inorder[p+1:], postorder[p:l-1])

	return root
}
