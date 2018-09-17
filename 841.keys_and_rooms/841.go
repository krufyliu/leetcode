package keys_and_rooms

import (
	"container/list"
)

func canVisitAllRooms(rooms [][]int) bool {
	steps := make([]bool, len(rooms))
	q := list.New()
	q.PushBack(0)
	for cur := q.Front(); cur != nil; cur = q.Front() {
		rn := cur.Value.(int)
		steps[rn] = true
		for _, nn := range rooms[rn] {
			if !steps[nn] {
				q.PushBack(nn)
			}
		}
		q.Remove(cur)
	}
	allIn := true
	for _, in := range steps {
		allIn = allIn && in
	}
	return allIn
}
