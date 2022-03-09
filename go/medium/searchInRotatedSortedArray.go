package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		pivot := (end+start)/2 
		if nums[pivot] == target {
			return pivot
		}

		if nums[start] <= nums[pivot] {
			if target < nums[start] || target > nums[pivot]{
				start = pivot + 1
			} else {
				end = pivot - 1
			}
		} else {
			if target > nums[end] || target < nums[pivot]{
				end = pivot -1
			} else {
				start = pivot +1
			}
		}
	}

	if nums[start] == target {
		return start
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 4))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 6))
	fmt.Println(search([]int{2,4, 5, 6, 7, 0, 1}, 7))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 1))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println("test")
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1}, 4))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1}, 5))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1}, 6))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1}, 7))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1}, 1))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1}, 2))

}

// 4,5,6,7,0,1,3 - 3 ->
