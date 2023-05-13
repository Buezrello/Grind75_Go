package solutions

func MaxProfit(prices []int) int {
	result := 0
	minBuy := prices[0]

	for _, price := range prices[1:] {
		if price < minBuy {
			minBuy = price
		}
		maxSell := price - minBuy
		if maxSell > result {
			result = maxSell
		}
	}

	return result
}
