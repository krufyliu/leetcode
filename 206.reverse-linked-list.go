/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre, cur, next *ListNode
	cur = head
	next = head.Next

	for next != nil {
		cur.Next = pre
		pre = cur
		cur = next
		next = cur.Next
	}

	cur.Next = pre

	return cur
}
