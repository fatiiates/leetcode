package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	quotient := 0
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sign := 1
	if !signsEqual(dividend, divisor) {
		sign = -1
	}

	dividend = abs(dividend)
	divisor = abs(divisor)

	for dividend-divisor >= 0 {
		shift := 0
		for dividend-(divisor << 1 << shift) >= 0 {
			shift++
		}
		dividend -= divisor << shift
		quotient += 1 << shift
	}

	return quotient * sign
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func signsEqual(n, m int) bool {
	if (n <= 0 && m <= 0) || (n >= 0 && m >= 0) {
		return true
	}
	return false
}

func main() {
	fmt.Println(divide(10, 2))
	fmt.Println(divide(math.MinInt32, -1))
	fmt.Println(divide(1000, 1))
	fmt.Println(divide(33, 5))
}
