package bestTimeToBuyAndSellStok2

func maxProfit(prices []int) int {

	var maxProfit int
	var buy, sell int
	var hold bool

	for i := 0; i < len(prices)-1; i++ {
		if !hold {
			if prices[i+1] > prices[i] {
				//fmt.Printf("buy: %d\n", prices[i])
				buy = prices[i]
				hold = true
			}
		}

		if hold {
			if prices[i+1] >= prices[i] {
				if i+1 == len(prices)-1 {
					sell = prices[i+1]
					hold = false
					maxProfit += sell - buy
				}
			} else {
				sell = prices[i]
				hold = false
				maxProfit += sell - buy
			}
		}

	}

	return maxProfit
}
