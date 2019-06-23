/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			name: "case 2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			name: "case 3",
			args: args{
				candidates: []int{1, 2},
				target:     4,
			},
			want: [][]int{
				{1, 1, 1, 1},
				{1, 1, 2},
				{2, 2},
			},
		},
		{
			name: "case 4",
			args: args{
				candidates: []int{1, 2, 3},
				target:     6,
			},
			want: [][]int{
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 2},
				{1, 1, 1, 3},
				{1, 1, 2, 2},
				{1, 2, 3},
				{2, 2, 2},
				{3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
