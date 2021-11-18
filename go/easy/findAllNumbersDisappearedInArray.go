package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	arr := []int{}
	l := len(nums)

	swapSort(&nums, l)

	for i := 0; i < l; i++ {
		if i+1 != nums[i] {
			arr = append(arr, i+1)
		}
	}
	return arr
}

func swapSort(nums *[]int, l int) {

	for i := 0; i < l; i++ {
		for i+1 != (*nums)[i] {
			(*nums)[i], (*nums)[(*nums)[i]-1] = (*nums)[(*nums)[i]-1], (*nums)[i]
			if (*nums)[i] > l || (*nums)[(*nums)[i]-1] == (*nums)[i] {
				break
			}
		}
	}

}

func main() {

	fmt.Println(findDisappearedNumbers([]int{4, 6, 3, 7, 2, 2, 3}))

}
