package leetcode

import (
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func TailList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	return tail
}

func (head *ListNode) nodeAt(idx int) *ListNode {
	cur := head
	for i := 0; i < idx; i++ {
		cur = cur.Next
	}
	return cur
}

func ListFromString(s string) *ListNode {
	if len(s) == 0 {
		return nil
	}
	ss := strings.Split(s, "->")

	val, err := strconv.Atoi(ss[0])
	if err != nil {
		panic(err)
	}

	head := &ListNode{
		Val: val,
	}

	cur := head

	for i := 1; i < len(ss); i++ {
		val, err := strconv.Atoi(ss[i])
		if err != nil {
			panic(err)
		}

		cur.Next = &ListNode{
			Val: val,
		}
		cur = cur.Next
	}

	return head
}

func (n *ListNode) String() string {
	var arr []string
	cur := n

	for cur != nil {
		arr = append(arr, strconv.Itoa(cur.Val))
		cur = cur.Next
	}

	return strings.Join(arr, "->")
}
