package main

import (
	"fmt"
)

func plusOne(digits []int) []int {

	d := len(digits) - 1

	carry := 0

	for ; d >= 0; d-- {
		if digits[d] == 9 {
			carry = 1
			digits[d] = 0
		} else {
			carry = 0
			digits[d] = digits[d] + 1
			break
		}
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{987}))
}
