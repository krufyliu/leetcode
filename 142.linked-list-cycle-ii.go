/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func detectCycle(head *ListNode) *ListNode {
	node := meetNode142(head)
	if node == nil {
		return nil
	}
	cs := circleSize(node)

	// 找到环入口
	p1, p2 := head, head
	for i := 0; i < cs; i++ {
		p2 = p2.Next
	}

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}

// 通过快慢指针找出是否有环，如果有就返回相遇的节点,
// 没有返回nil
func meetNode142(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pSlow, pFast := head, head.Next

	for pFast != nil {
		if pSlow == pFast {
			return pFast
		}
		pSlow = pSlow.Next

		// 快指针一次移动两位
		pFast = pFast.Next
		if pFast == nil {
			return nil
		}
		pFast = pFast.Next
	}

	return nil
}

func circleSize(head *ListNode) int {
	size := 1
	next := head.Next
	for next != head {
		size++
		next = next.Next
	}

	return size
}
