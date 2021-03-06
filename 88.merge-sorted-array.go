package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums2[j] >= nums1[i] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}

	if j >= 0 {
		for j >= 0 {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}
}
