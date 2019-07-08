/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}

	var paths []string
	if root.Left != nil {
		subPaths := binaryTreePaths(root.Left)
		for _, path := range subPaths {
			paths = append(paths, fmt.Sprintf("%d->%s", root.Val, path))
		}
	}

	if root.Right != nil {
		subPaths := binaryTreePaths(root.Right)
		for _, path := range subPaths {
			paths = append(paths, fmt.Sprintf("%d->%s", root.Val, path))
		}
	}

	return paths
}
