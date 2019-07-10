/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
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

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case 1",
			args: args{
				head: ListFromString("1->2->3->4"),
			},
			want: ListFromString("2->1->4->3"),
		},
		{
			name: "case 2",
			args: args{
				head: ListFromString("1->2->3"),
			},
			want: ListFromString("2->1->3"),
		},
		{
			name: "case 3",
			args: args{
				head: ListFromString("1"),
			},
			want: ListFromString("1"),
		},
		{
			name: "case 4",
			args: args{
				head: ListFromString(""),
			},
			want: ListFromString(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); got.String() != tt.want.String() {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
