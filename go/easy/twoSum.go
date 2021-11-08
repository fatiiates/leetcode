package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if i == j {
				continue
			}
			fmt.Println(i, j)
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func main() {

	fmt.Println(twoSum([]int{3, 2, 4}, 6))

}
