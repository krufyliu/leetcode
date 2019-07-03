/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var preReverse *ListNode
	cur := head
	for i := 1; i < m; i++ {
		preReverse = cur
		cur = cur.Next
	}

	rangeStart := cur
	rangeEnd := cur
	var rangeStartPre *ListNode
	var rangeEndNext *ListNode
	for i := m; i <= n; i++ {
		rangeEnd = cur
		rangeEndNext = cur.Next
		cur.Next = rangeStartPre
		rangeStartPre = cur
		cur = rangeEndNext
	}

	if preReverse != nil {
		preReverse.Next = rangeEnd
		rangeStart.Next = rangeEndNext
		return head
	}

	if rangeEndNext != nil {
		rangeStart.Next = rangeEndNext
	}

	return rangeEnd

}
