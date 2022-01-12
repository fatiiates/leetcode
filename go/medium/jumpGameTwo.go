package main

func jump(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 10001
	}
	dp[l-1] = 0

	for i := l - 2; i >= 0; i-- {
		for j := 1; j <= nums[i]; j++ {
			dp[i] = min(dp[i], 1+dp[min(l-1, i+j)])
		}
	}
	return dp[0]
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
