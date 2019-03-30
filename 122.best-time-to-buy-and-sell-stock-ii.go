package leetcode

func maxProfit(prices []int) int {
	maxProfit := 0
	l := len(prices)

	if l < 2 {
		return maxProfit
	}

	for i := 0; i < l-1; i++ {
		if prices[i] >= prices[i+1] {
			continue
		}
		var j int
		for j = i + 1; j < l; j++ {
			if prices[j] < prices[j-1] {
				maxProfit += prices[j-1] - prices[i]
				break
			}
		}
		if j == l {
			maxProfit += prices[j-1] - prices[i]
		}
		i = j - 1
	}

	return maxProfit
}
