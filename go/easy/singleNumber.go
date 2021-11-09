package main

import (
	"fmt"
)

func singleNumber(nums []int) int {

	var arr map[int]int = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		fmt.Println(arr)
		if _, ok := arr[nums[i]]; ok {
			delete(arr, nums[i])
		} else {
			arr[nums[i]] = 1
		}
	}

	for k, _ := range arr {
		return k
	}
	return 0
}

func main() {

	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))

}
