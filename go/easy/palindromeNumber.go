package main

import (
	"fmt"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	nums := []int{}
	for ; x > 0; x /= 10 {
		nums = append(nums, x%10)
	}

	l := len(nums) - 1
	for i, v := range nums {
		if v != nums[l-i] {
			return false
		}
	}

	return true
}

func main() {

	fmt.Println(isPalindrome(-1231))

}
