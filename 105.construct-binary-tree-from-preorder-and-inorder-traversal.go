/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (39.38%)
 * Total Accepted:    208K
 * Total Submissions: 520.1K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * Given preorder and inorder traversal of a tree, construct the binary tree.
 *
 * Note:
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * preorder = [3,9,20,15,7]
 * inorder = [9,3,15,20,7]
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

func buildTree105(preorder []int, inorder []int) *TreeNode {
	l1 := len(preorder)

	if l1 == 0 {
		return nil
	}

	first := preorder[0]
	root := &TreeNode{
		Val: first,
	}

	if l1 == 1 {
		return root
	}

	var p int
	for i, v := range inorder {
		if v == first {
			p = i
			break
		}
	}
	root.Left = buildTree(preorder[1:p+1], inorder[0:p])
	root.Right = buildTree(preorder[p+1:], inorder[p+1:])

	return root
}
