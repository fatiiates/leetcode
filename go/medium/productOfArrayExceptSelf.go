package main

import "fmt"

func productExceptSelf(nums []int) []int {

	l := len(nums)
	arr := make([]int, l)
	arr[0] = 1
	for i := 1; i < l; i++ {
		arr[i] = arr[i-1] * nums[i-1]
	}

	total := nums[l-1]
	for i := l - 2; i >= 0; i-- {
		arr[i] *= total
		total *= nums[i]
	}

	return arr
}

func main() {

	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
