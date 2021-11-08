package main

import "fmt"

func maxProfit(prices []int) int {

	mn := 100001
	mx := 0

	for i := 0; i < len(prices); i++ {

		mn = min(mn, prices[i])
		mx = max(mx, prices[i]-mn)

	}

	return mx
}

func min(n1 int, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}
func max(n1 int, n2 int) int {
	if n1 < n2 {
		return n2
	}
	return n1
}

func main() {

	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))

}
