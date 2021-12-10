package main

import "fmt"

var mod int = 1e9 + 7

func numTilings(n int) int {
	if n <= 2 {
		return n
	}

	filled_prev, gap_prev, filled_prev2, gap_prev2 := 2, 2, 1, 1
	for i := 3; i <= n; i++ {
		filled := (filled_prev + filled_prev2 + 2*gap_prev2) % mod
		gap := (filled_prev + gap_prev) % mod
		filled_prev2, filled_prev = filled_prev, filled
		gap_prev2, gap_prev = gap_prev, gap
	}

	return filled_prev
}

func main() {

	fmt.Println(numTilings(20))
}
