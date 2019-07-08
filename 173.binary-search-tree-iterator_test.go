/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
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

import (
	"testing"
)

func TestBSTIterator(t *testing.T) {
	bt := NewBinaryTreeFromString("7,3,15,null,null,9,20")
	iter := Constructor(bt)
	sortArr := []int{3, 7, 9, 15, 20}
	for i := 0; i < 5; i++ {
		if !iter.HasNext() {
			t.Errorf("call next shoud be true")
		}
		n := iter.Next()
		if n != sortArr[i] {
			t.Errorf("call next want get %d, but got %d", sortArr[i], n)
		}
	}

	if iter.HasNext() {
		t.Errorf("iterator should be no more")
	}
}
