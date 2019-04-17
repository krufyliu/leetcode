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

func ListFromString(s string) *ListNode {
	ss := strings.Split(s, "->")
	if len(ss) == 0 {
		return nil
	}

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
