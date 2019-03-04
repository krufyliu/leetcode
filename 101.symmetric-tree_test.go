package leetcode

import (
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				root: NewBinaryTreeFromString("1,2,2,3,4,4,3"),
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				root: NewBinaryTreeFromString("1,2,2,null,3,null,3"),
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				root: NewBinaryTreeFromString("1,2,2,3,null,null,3"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSymmetricRecursively(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				root: NewBinaryTreeFromString("1,2,2,3,4,4,3"),
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				root: NewBinaryTreeFromString("1,2,2,null,3,null,3"),
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				root: NewBinaryTreeFromString("1,2,2,3,null,null,3"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetricRecursively(tt.args.root); got != tt.want {
				t.Errorf("isSymmetricRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}
