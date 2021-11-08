package main

import (
	"fmt"
	"math"
)

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, 0)
}

func binarySearch(nums []int, target int, start int) int {

	med := int(math.Floor(float64(len(nums)) / 2))

	if len(nums) == 0 {
		return start
	}
	if nums[med] == target {
		return start + med
	}
	if target < nums[med] {
		return binarySearch(nums[:med], target, start)
	}
	if target > nums[med] {
		return binarySearch(nums[med+1:], target, start+med+1)
	}
	return start
}

func main() {
	fmt.Println(searchInsert([]int{2}, 1))
}
