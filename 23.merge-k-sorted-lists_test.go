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
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case 1",
			args: args{
				lists: []*ListNode{
					ListFromString("1->4->5"),
					ListFromString("1->3->4"),
					ListFromString("2->6"),
				},
			},
			want: ListFromString("1->1->2->3->4->4->5->6"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeKLists(tt.args.lists)
			t.Logf("%s", got.String())
			if got.String() != tt.want.String() {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
