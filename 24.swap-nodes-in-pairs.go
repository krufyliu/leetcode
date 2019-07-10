/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead, last *ListNode
	cur, next := head, head.Next
	for {
		if newHead == nil {
			newHead = next
		}

		next.Next, cur.Next = cur, next.Next
		if last != nil {
			last.Next = next
		}
		last = cur

		cur = cur.Next
		if cur == nil || cur.Next == nil {
			break
		}

		next = cur.Next
	}

	return newHead
}
