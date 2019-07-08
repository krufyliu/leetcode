/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package leetcode

import (
	"container/heap"
)

type mergeHeap []*ListNode

func (h mergeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h mergeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h mergeHeap) Len() int {
	return len(h)
}

func (h *mergeHeap) Pop() interface{} {
	l := h.Len()
	var p interface{}
	if l > 0 {
		p = (*h)[l-1]
		*h = (*h)[0 : l-1]
	}
	return p
}

func (h *mergeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &mergeHeap{}

	var mergeHead, cur, pre, next *ListNode
	for _, head := range lists {
		if head != nil {
			*h = append(*h, head)
		}
	}
	heap.Init(h)

	for h.Len() > 0 {
		node := heap.Pop(h)
		cur = node.(*ListNode)
		if mergeHead == nil {
			mergeHead = cur
			pre = mergeHead
		} else {
			pre.Next = cur
			pre = cur
		}

		next = cur.Next
		if next != nil {
			heap.Push(h, next)
		}
	}

	return mergeHead
}
