package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	l1 := len(nums1)
	l2 := len(nums2)
	stack := make([]int, l2)
	solution := make([]int, l1)
	next := make(map[int]int)

	sp := -1
	for i := 0; i < l2; i++ {
		for sp >= 0 && stack[sp] < nums2[i] {
			next[stack[sp]] = nums2[i]
			sp--
		}
		sp++
		stack[sp] = nums2[i]
	}

	for sp >= 0 {
		next[stack[sp]] = -1
		sp--
	}

	for i := 0; i < l1; i++ {
		solution[i] = next[nums1[i]]
	}

	return solution
}
