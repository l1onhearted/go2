package main

//买卖股票最佳时机
func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-2;i++ {
		if prices[i+1] > prices[i] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}
