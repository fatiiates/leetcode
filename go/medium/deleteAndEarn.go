package main

import "fmt"

func deleteAndEarn(nums []int) int {
	l := len(nums)
	res := make([]int, 10001)
	mx := 0
	for i := 0; i < l; i++ {
		res[nums[i]] += nums[i]
		mx = max(mx, nums[i])
	}

	dp := make([]int, mx+1)
	if l == 1 {
		return nums[0]
	}
	dp[1] = res[1]
	for i := 2; i < l; i++ {
		dp[i] = max(dp[i-2]+res[i], dp[i-1])
	}
	fmt.Println(dp[:10], mx)
	return dp[mx]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	fmt.Println(deleteAndEarn([]int{8, 2, 1, 2, 5, 3, 3, 3, 4, 9,15}))
}