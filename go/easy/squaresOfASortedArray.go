package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {

	s := 0
	e := len(nums) - 1

	arr := []int{}

	for s <= e {

		if abs(nums[s]) >= abs(nums[e]) {
			arr = append(arr, nums[s]*nums[s])
			s++
		} else {
			arr = append(arr, nums[e]*nums[e])
			e--
		}

	}

	reverse(&arr)

	return arr

}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func reverse(arr *[]int) {
	l := len(*arr)
	for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func main() {

	fmt.Println(sortedSquares([]int{-5, -4, -3, -2, -1}))

}
