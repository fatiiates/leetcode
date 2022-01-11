package main

import "fmt"

var mx int = 0

func maxSubArray(nums []int) int {
	mx = nums[0]
	return helper(&nums, 0, len(nums)-1)
}

func helper(nums *[]int, l, r int) int {
	if l == r {
		return (*nums)[l]
	}

	p := (r - l) / 2
	sLeft := helper(nums, l, p)
	sRight := helper(nums, p+1, r)

	sLeft = max(sLeft+sRight, sRight)
	mx = max(sLeft, mx)

	return mx
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
