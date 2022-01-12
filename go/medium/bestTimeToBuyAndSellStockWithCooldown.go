package main

import "fmt"

func maxProfit(prices []int) int {
	l := len(prices)
	if l == 1 {
		return 0
	}
	buy := max(-prices[0], -prices[1])
	prev := 0
	sell := max(0, buy+prices[1])
	for i := 2; i < l; i++ {
		buy = max(buy, prev-prices[i])  
		prev = sell
		sell = max(sell, buy+prices[i])
	}

	return sell

}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {

	fmt.Println(maxProfit([]int{1, 2, 3, 2}))
}
