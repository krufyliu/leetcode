package jewels_and_stones

import (
	"strings"
)

func numJewelsInStones(J string, S string) int {
	var count = 0
	for _, c := range S {
		if strings.IndexRune(J, c) >= 0 {
			count++
		}
	}
	return count
}
