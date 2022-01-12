package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	l := len(prices)
	cash, hold := 0, -prices[0]
	for i := 1; i < l; i++ {
		cash = max(cash, hold+prices[i]-fee)
		hold = max(hold, cash-prices[i])
	}

	return cash
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {

	fmt.Println(maxProfit([]int{1, 6, 2, 8, 4, 9}, 2))
}
