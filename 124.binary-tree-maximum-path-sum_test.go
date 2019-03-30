package leetcode

import "testing"

func Test_maxPathSum(t *testing.T) {
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
				root: NewBinaryTreeFromString("1,2,3"),
			},
			want: 6,
		},
		{
			name: "case 2",
			args: args{
				root: NewBinaryTreeFromString("-10,9,20,null,null,15,7"),
			},
			want: 42,
		},
		{
			name: "case 3",
			args: args{
				root: NewBinaryTreeFromString("-5,2,3,4,1,null,null,-1"),
			},
			want: 7,
		},
		{
			name: "case 4",
			args: args{
				root: NewBinaryTreeFromString("-1,-2,-3"),
			},
			want: -1,
		},
		{
			name: "case 5",
			args: args{
				root: NewBinaryTreeFromString("1,-2,-3,1,3,-2,null,-1"),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
