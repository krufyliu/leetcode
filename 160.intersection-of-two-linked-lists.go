/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p1, p2 := headA, headB
	var tail1, tail2 *ListNode

	for p1 != p2 {
		if p1.Next != nil {
			p1 = p1.Next
		} else {
			if tail1 == nil {
				tail1 = p1
			}
			p1 = headB
		}

		if p2.Next != nil {
			p2 = p2.Next
		} else {
			if tail2 == nil {
				tail2 = p2
			}
			p2 = headA
		}

		if tail1 != nil && tail2 != nil && tail1 != tail2 {
			return nil
		}
	}

	return p1
}

func meetNode(head *ListNode) *ListNode {
	pslow, pFast := head, head.Next

	for pFast != nil {
		if pslow == pFast {
			return pFast
		}

		pslow = pslow.Next
		pFast = pFast.Next
		if pFast == nil {
			return nil
		}
		pFast = pFast.Next
	}

	return pFast
}

func circleNodeCount(head *ListNode) int {
	next := head.Next
	count := 1
	for next != head {
		next = next.Next
		count++
	}

	return count
}

func circleEntry(head *ListNode) *ListNode {
	mn := meetNode(head)
	if mn == nil {
		return nil
	}
	cn := circleNodeCount(mn)

	p1, p2 := head, head
	// p1先移动cn步
	for i := 1; i <= cn; i++ {
		p1 = p1.Next
	}

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}
