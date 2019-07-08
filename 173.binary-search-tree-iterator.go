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

import "container/list"

type BSTIterator struct {
	stack *list.List
}

func Constructor(root *TreeNode) BSTIterator {
	li := list.New()
	iter := BSTIterator{stack: li}
	iter.add(root)
	return iter
}

func (this *BSTIterator) add(node *TreeNode) {
	for node != nil {
		this.stack.PushBack(node)
		node = node.Left
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	ele := this.stack.Back()
	this.stack.Remove(ele)
	node := ele.Value.(*TreeNode)
	// 右子树入栈
	if node.Right != nil {
		this.add(node.Right)
	}
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.stack.Len() != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
