package main

import (
	"fmt"
)

func addBinary(a string, b string) string {

	max := a
	min := b
	if len(a) < len(b) {
		max = b
		min = a
	}

	maxLen := len(max)
	minLen := len(min)

	for i := 0; i < maxLen-minLen; i++ {
		min = "0" + min
	}
	minLen = len(min)

	sum := ""
	carry := 0
	for i := 1; i <= maxLen; i++ {
		if max[maxLen-i] == min[minLen-i] {
			if max[maxLen-i] == 49 {
				if carry == 1 {
					sum = "1" + sum
				} else {
					sum = "0" + sum
				}
				carry = 1
			} else {
				if carry == 1 {
					sum = "1" + sum
					carry = 0
				} else {
					sum = "0" + sum
				}
			}
		} else {
			if max[maxLen-i] == 49 || min[minLen-i] == 49 {
				if carry == 1 {
					carry = 1
					sum = "0" + sum
				} else {
					sum = "1" + sum
				}
			} else {
				if carry == 1 {
					sum = "1" + sum
					carry = 0
				} else {
					sum = "0" + sum
				}
			}
		}

	}
	if carry == 1 {
		sum = "1" + sum
	}
	return sum
}

func main() {
	fmt.Println(addBinary("1010", "1011"))
}
