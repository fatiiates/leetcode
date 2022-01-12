package main

import "math"

func maxSubarraySumCircular(nums []int) int {
	S := 0
	for _, v := range nums {
		S += v
	}

	if len(nums) == 1 {
		return S
	}

	case_1 := math.MinInt
	cur := math.MinInt
	for _, x := range nums {
		cur = x + max(cur, 0)
		case_1 = max(case_1, cur)
	}

	case_2 := math.MaxInt
	cur = math.MaxInt
	for i := 1; i < len(nums); i++ {
		cur = nums[i] + min(cur, 0)
		case_2 = min(case_2, cur)
	}
	case_2 = S - case_2

	case_3 := math.MaxInt
	cur = math.MaxInt
	for i := 0; i < len(nums)-1; i++ {
		cur = nums[i] + min(cur, 0)
		case_3 = min(case_3, cur)
	}

	return max(case_3, max(case_3, case_3))
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
