/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
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

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case 1",
			args: args{
				head: ListFromString("1->2->3->4->5"),
				m:    2,
				n:    4,
			},
			want: ListFromString("1->4->3->2->5"),
		},
		{
			name: "case 2",
			args: args{
				head: ListFromString("1->2->3->4"),
				m:    2,
				n:    4,
			},
			want: ListFromString("1->4->3->2"),
		},
		{
			name: "case 3",
			args: args{
				head: ListFromString("1->2->3->4->5"),
				m:    1,
				n:    4,
			},
			want: ListFromString("4->3->2->1->5"),
		},
		{
			name: "case 4",
			args: args{
				head: ListFromString("1->2->3->4->5"),
				m:    1,
				n:    5,
			},
			want: ListFromString("5->4->3->2->1"),
		},
		{
			name: "case 5",
			args: args{
				head: ListFromString("1->2->3"),
				m:    3,
				n:    3,
			},
			want: ListFromString("1->2->3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
