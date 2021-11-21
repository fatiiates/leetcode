package main

import (
	"fmt"
)

func singleNumber(nums []int) int {

	r := nums[0]
	l := len(nums)
	for i := 1; i < l; i++ {
		r ^= nums[i]
	}
	return r
}

func main() {

	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))

}
