/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.13%)
 * Total Accepted:    376.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 *
 * Example:
 *
 *
 * Given linked list: 1->2->3->4->5, and n = 2.
 *
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 *
 *
 * Note:
 *
 * Given n will always be valid.
 *
 * Follow up:
 *
 * Could you do this in one pass?
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	for i := 1; i < n; i++ {
		cur = cur.Next
	}

	if cur.Next == nil {
		head = head.Next
		return head
	}
	// k+1
	cur = cur.Next
	target := head
	for cur.Next != nil {
		target = target.Next
		cur = cur.Next
	}

	target.Next = target.Next.Next
	return head
}
