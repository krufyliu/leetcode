package removeduplicatesfromsortedarray

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{1, 1, 1, 1, 1, 1},
			},
			want: 1,
		},
		{
			name: "case 6",
			args: args{
				nums: []int{1, 1, 2, 2, 3, 3},
			},
			want: 3,
		},
		{
			name: "case 7",
			args: args{
				nums: []int{1, 1, 2, 3, 4, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
