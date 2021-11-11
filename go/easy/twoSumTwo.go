package main

import (
	"fmt"
	"math"
)

func twoSum(nums []int, target int) []int {

	nStart := search(&nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {

		if target > nums[i] {
			index := 0
			j := 0
			if i >= nStart+1 {
				j = i + 1
				index = binarySearch(nums[i+1:], target-nums[i], 0)
			} else {
				j = nStart + 1
				index = binarySearch(nums[nStart+1:], target-nums[i], 0)
			}

			if index != -1 {
				return []int{i + 1, index + j + 1}
			}
		} else if target == nums[i] {
			if nums[nStart+1] == 0 {
				return []int{i + 1, nStart + 2}
			}
		} else {
			index := binarySearch(nums[i+1:nStart+1], target-nums[i], 0)
			if index != -1 {
				return []int{i + 1, index + i + 2}
			}
		}

	}

	return []int{}
}

func isNonNegative(version int) bool {
	if version >= 0 {
		return true
	}
	return false
}

func search(nums *[]int, start int, last int) int {

	med := int(math.Ceil(float64((last - start)) / 2.0))

	res := isNonNegative((*nums)[start+med])

	if last == start {
		return start
	}
	if res {
		return search(nums, start, last-med)
	} else {
		return search(nums, start+med, last)
	}

}

func binarySearch(nums []int, target int, start int) int {
	med := int(math.Floor(float64(len(nums)) / 2))

	if len(nums) == 0 {
		return -1
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

	fmt.Println(twoSum([]int{-5, -4, -3, 0}, -5))

}
