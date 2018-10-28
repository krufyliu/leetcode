package heap

import (
	"log"
	"testing"
)

func Test_heapfiy(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}

	arr := []int{4, 5, 6, 1, 2, 3, 8, 9, 10}
	heapfiy(arr)
	log.Printf("%#v", arr)

	arr[0] = 11
	down(arr, 0, len(arr))
	log.Printf("%#v", arr)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapfiy(tt.args.arr)
		})
	}
}
