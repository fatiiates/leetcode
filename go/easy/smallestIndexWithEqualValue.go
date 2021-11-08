package main

import (
	"fmt"
)

func smallestEqual(nums []int) int {
	smallest := 101
	for i, v := range nums {
		x := i % 10

		if x == v && i < smallest {
			smallest = i
		}
	}

	if smallest == 101 {
		return -1
	}
	return smallest
}

func main() {
	fmt.Println(smallestEqual([]int{7, 8, 3, 5, 2, 6, 3, 1, 1, 4, 5, 4, 8, 7, 2, 0, 9, 9, 0, 5, 7, 1, 6}))
}
