package main

import (
	"fmt"
)

func singleNumber(nums []int) []int {

	var list map[int]int = make(map[int]int)
	arr := []int{}

	for i := 0; i < len(nums); i++ {
		if _, ok := list[nums[i]]; ok {
			list[nums[i]] += 1
		} else {
			list[nums[i]] = 1
		}
	}

	for k, v := range list {
		if v == 1 {
			arr = append(arr, k)
		}
	}
	return arr
}

func main() {

	fmt.Println(singleNumber([]int{-1, 0}))
}
