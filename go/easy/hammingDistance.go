package main

import (
	"fmt"
)

func hammingDistance(x int, y int) int {

	if x == y {
		return 0
	}
	mx, mn := max(x, y)
	dist := 0
	for mx > 0 {
		d1 := mx % 2
		d2 := mn % 2
		if d1 != d2 {
			dist++
		}
		mn /= 2
		mx /= 2
	}
	fmt.Println(mn, mx)

	return dist
}

func max(n1 int, n2 int) (int, int) {
	if n1 > n2 {
		return n1, n2
	}
	return n2, n1
}

func main() {

	fmt.Println(hammingDistance(5, 8))

}
