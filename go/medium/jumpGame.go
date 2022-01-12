package main

import "fmt"

func canJump(nums []int) bool {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = nums[0]
	if dp[0] == 0 && l > 1 {
		return false
	}
	for i := 1; i < l-1; i++ {
		dp[i] = max(dp[i-1]-1, nums[i])

		if dp[i] == 0 {
			return false
		}
	}
	fmt.Println(dp)
	return true
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	fmt.Println(canJump([]int{1, 2, 3}))
	fmt.Println(canJump([]int{2, 0}))
	fmt.Println(canJump([]int{1}))
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{0, 1}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{3, 1, 0, 4}))
}
