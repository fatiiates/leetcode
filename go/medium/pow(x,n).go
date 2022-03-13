package main

import "fmt"

func myPow(x float64, n int) float64 {
	var res float64 = 1
	ni := abs(n)
	for ni > 0 {
		// n is odd
		if ni&1 == 1 {
			res *= x
		}

		ni >>= 1 //to divide n by 2
		x *= x
	}
	if n < 0 {
		return 1 / res
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(myPow(2, -2147483648))
}
