package main

import "fmt"

func arrangeCoins(n int) int {

	total := 0
	stairs := 0
	for i := 1; total <= n; i, stairs = i+1, stairs+1 {
		total += i
	}

	if total > n {
		return stairs - 1
	}
	return stairs
}

func main() {

	fmt.Println(arrangeCoins(12))

}
