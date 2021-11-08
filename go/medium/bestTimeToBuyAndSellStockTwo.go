package main

import "fmt"

func maxProfit(prices []int) int {

	return calcProfit(&prices, 0, 0)

}

func calcProfit(prices *[]int, index int, total int) int {
	mn := 100001
	mx := 0

	for i := index; i < len(*prices); i++ {

		mn = min(mn, (*prices)[i])
		if (*prices)[i]-mn > mx {
			if i != len(*prices)-1 && (*prices)[i+1] <= (*prices)[i] {
				total += (*prices)[i] - mn + calcProfit(prices, i+1, total)
				break
			}
			mx = (*prices)[i] - mn

		}

	}

	if total == 0 {
		return mx
	}
	return total

}

func min(n1 int, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func main() {

	fmt.Println(maxProfit([]int{5, 2, 3, 2, 6, 6, 2, 9, 1, 0, 7, 4, 5, 0}))
}
