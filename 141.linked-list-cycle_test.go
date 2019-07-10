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

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	head1 := ListFromString("3->2->0->-4")
	n3 := head1.nodeAt(3)
	n1 := head1.nodeAt(1)
	n3.Next = n1
	head2 := ListFromString("1->2")
	head2.nodeAt(1).Next = head2.nodeAt(0)
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				head: head1,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				head: head2,
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				head: ListFromString("1"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
