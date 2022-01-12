package main

func minCostClimbingStairs(cost []int) int {

	dp := make([]int, 1001)
	l := len(cost)
	for i := 2; i <= l; i++ {
		dp[i] += min(dp[i-2] + cost[i-2], dp[i-1] + cost[i-1])
	}
	return dp[l]

}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
