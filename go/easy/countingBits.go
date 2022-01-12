package main

func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = (i & 1) + dp[i/2]
	}
	return dp
}
