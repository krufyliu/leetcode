package leetcode

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				root: NewBinaryTreeFromString("3,9,20,null,null,15,7"),
			},
			want: [][]int{
				{15, 7},
				{9, 20},
				{3},
			},
		},
		{
			name: "case 2",
			args: args{
				root: NewBinaryTreeFromString("1,2,null,3,null,4,null"),
			},
			want: [][]int{
				{4},
				{3},
				{2},
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
