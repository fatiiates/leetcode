package main

import "fmt"

func searchRange(nums []int, target int) []int {
	arr := []int{-1, -1}

	start := 0
	end := len(nums) - 1

	// first occurence
	for start < end {
		p := (start + end) / 2

		if target > nums[p] {
			start = p + 1
		} else if target == nums[p] {
			end = p
		} else {
			end = p - 1
		}
	}
	if start < len(nums) && nums[start] == target {
		arr[0] = start
	}
	start = 0
	end = len(nums) - 1

	// last occurence
	for start <= end {
		p := (start + end) / 2

		if target < nums[p] {
			end = p - 1
		} else {
			start = p + 1
		}
	}

	if end >= 0 && nums[end] == target {
		arr[1] = end
	}

	return arr
}

func main() {
	fmt.Println(searchRange([]int{5, 8, 8, 8, 8, 8, 8, 8, 8}, 8))
	fmt.Println(searchRange([]int{8}, 8))
	fmt.Println(searchRange([]int{8}, 0))
}
