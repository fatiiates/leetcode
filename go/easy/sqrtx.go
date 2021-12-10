package main

import (
	"fmt"
)

func mySqrt(x int) int {

	y := primeMultiplier(x)
	return y
}

func primeMultiplier(x int) int {

	if x <= 0 {
		return 0
	}
	for i := 2; ; i++ {
		if i*i > x {
			return i - 1
		}
		if i*i == x {
			return i
		}
	}
}

func main() {
	fmt.Println(mySqrt(28))
}
