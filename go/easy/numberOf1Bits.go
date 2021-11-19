package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {

	if num == 0 {
		return 0
	}

	dist := 0
	for ; num > 0; num = num >> 1 {
		if num&1 == 1 {
			dist += 1
		}
	}
	return dist
}

func main() {

	fmt.Println(hammingWeight(3))
}
