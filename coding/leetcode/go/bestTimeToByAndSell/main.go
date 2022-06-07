package bestTimeToByAndSell

import "fmt"

func maxProfit(prices []int) int {
	var maxP, sell, buy int

	for buy < len(prices) && sell < len(prices)-1 {

		sell++

		if prices[buy] > prices[sell] {
			buy = sell
			continue
		}

		fmt.Printf("buy: %d, sell: %d\n", buy, sell)

		newP := prices[sell] - prices[buy]

		fmt.Printf("newP: %d\n", newP)

		if newP < 0 {
			buy++
		} else {
			if maxP >= newP {
				continue
			} else {
				maxP = newP
			}
		}
	}
	return maxP
}
