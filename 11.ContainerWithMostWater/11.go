package ContainerWithMostWater

func maxArea(height []int) int {
	var max = 0
	l := len(height)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			mh := height[i]
			if height[i] > height[j] {
				mh = height[j]
			}
			area := mh * (j - i)
			if area > max {
				max = area
			}
		}
	}

	return max
}
