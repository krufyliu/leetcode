package leetcode

import "testing"

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				root: NewBinaryTreeFromString("3,9,20,null,null,15,7"),
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				root: NewBinaryTreeFromString(""),
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				root: NewBinaryTreeFromString("3,9,null,15,7"),
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				root: NewBinaryTreeFromString("3,9"),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
