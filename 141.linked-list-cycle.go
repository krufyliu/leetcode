/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	pSlow, pFast := head, head.Next

	for pFast != nil {
		if pSlow == pFast {
			return true
		}
		pSlow = pSlow.Next

		// 快指针一次移动两位
		pFast = pFast.Next
		if pFast == nil {
			return false
		}
		pFast = pFast.Next
	}

	return false
}
