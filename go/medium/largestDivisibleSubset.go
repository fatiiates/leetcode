package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {

	l := len(nums)
	dp := make([]int, l)
	pre := make([]int, l)
	res := []int{}

	sort.Ints(nums)

	ml := 0
	end := 0

	for i := 0; i < l; i++ {
		c := 1
		prei := -1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[j]+1 > c {
				c = dp[j] + 1
				prei = j
			}
		}
		dp[i] = c
		pre[i] = prei

		if dp[i] > ml {
			ml = dp[i]
			end = i
		}
	}

	for end != -1 {
		res = append(res, nums[end])
		end = pre[end]
	}

	return res

}

func main() {

	fmt.Println(largestDivisibleSubset([]int{4, 8, 10, 240}))
}
