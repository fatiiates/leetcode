package main

import "fmt"

var total int

func climbStairs(n int) int {
	var (
		step1 = 1
		step2 = 2
	)

	switch n {
	case 1:
		return step1
	case 2:
		return step2
	}

	res := 0

	for i := 2; i < n; i++ {
		res = step1 + step2
		step1 = step2
		step2 = res
	}

	return res
}

func main() {
	fmt.Println(climbStairs(45))
}
