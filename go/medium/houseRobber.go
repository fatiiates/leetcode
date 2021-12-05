package main

import "fmt"

func rob(nums []int) int {

	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	a := nums[0]
	b := max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		m := max(b, a+nums[i])
		a = b
		b = m
	}

	return b
}

func max(n1 int, n2 int) int {
	if n1 < n2 {
		return n2
	}
	return n1
}

func main() {

	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
