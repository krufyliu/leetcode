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

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				head: ListFromString("1->2->3->4->5"),
				n:    2,
			},
			want: "1->2->3->5",
		},
		{
			name: "case 2",
			args: args{
				head: ListFromString("1->2->3->4->5"),
				n:    5,
			},
			want: "2->3->4->5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Errorf(tt.args.head.String())
			if got := removeNthFromEnd(tt.args.head, tt.args.n); got.String() != tt.want {
				t.Errorf("removeNthFromEnd() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}
