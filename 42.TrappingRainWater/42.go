package TrappingRainWater

func trap(height []int) int {
	var units = 0
	l := len(height)

	var i, j, k, m int
	for i = 0; i < l-1; i++ {
		if height[i+1] < height[i] {
			break
		}
	}

	for k = l - 1; k > 0; k-- {
		if height[k-1] < height[k] {
			break
		}
	}

	var max = 0
	var maxi = i

	for m = i; m < l; m++ {
		if height[m] > max {
			max = height[m]
			maxi = m
		}
	}

	for i < maxi-1 {
		for j = i + 1; j <= maxi; j++ {
			if height[i] <= height[j] && j-i > 1 {
				for m = i + 1; m < j; m++ {
					units += (height[i] - height[m])
				}
				break
			} else {
				if height[i] > height[j] {
					continue
				} else {
					break
				}
			}
		}
		i = j
	}

	for k > maxi+1 {
		for j = k - 1; j >= maxi; j-- {
			if height[j] >= height[k] && k-j > 1 {
				for m = k - 1; m > j; m-- {
					units += (height[k] - height[m])
				}
				break
			} else {
				if height[k] > height[j] {
					continue
				} else {
					break
				}
			}
		}
		k = j
	}

	return units
}
