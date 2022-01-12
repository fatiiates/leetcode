package main

import "fmt"

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}
	mx := 0
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i < n; i++ {
		if n&1 == 0 {
			dp[i] = dp[i/2]
			mx = max(mx, dp[i])
		} else {
			dp[i+1] = dp[i/2] + dp[i/2+1]
			mx = max(mx, dp[i+1])
		}
	}

	return mx
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	fmt.Println(getMaximumGenerated(100))
}
