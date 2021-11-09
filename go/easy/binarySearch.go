package main

import (
	"fmt"
	"math"
)

func search(nums []int, target int) int {

	med := int(math.Floor(float64(len(nums)) / 2))

	if nums[med] == target {
		return med
	} else if len(nums) == 1 {
		return -1
	}
	if target < nums[med] {
		return search(nums[:med], target)
	} else if target > nums[med] {
		n := search(nums[med:], target)
		if n != -1 {
			return med + n
		} else {
			return -1
		}

	}

	return 0
}

func main() {

	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))

	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
