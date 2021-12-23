package main

func canPartition(nums []int) bool {
	sum := 0

	for _, i := range nums {
		sum += i
	}

	if sum&1 == 1 {
		return false
	}

	sz := sum >> 1

	dp := make([]bool, sz+1)
	dp[0] = true

	for _, i := range nums {
		for j := sz; j >= i; j-- {
			if dp[j-i] {
				dp[j] = true
			}
		}
	}
	return dp[sz]
}
