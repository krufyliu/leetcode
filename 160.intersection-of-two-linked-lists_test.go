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

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	headCommon := ListFromString("8->4->5")
	head1 := ListFromString("4->1")
	head2 := ListFromString("5->0->1")
	tail1 := TailList(head1)
	tail2 := TailList(head2)
	tail1.Next = headCommon
	tail2.Next = headCommon
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				headA: head1,
				headB: head2,
			},
			want: headCommon,
		},
		{
			name: "case 1",
			args: args{
				headA: ListFromString("2->6->4"),
				headB: ListFromString("1->5"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
