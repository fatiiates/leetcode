package main

func maxSubArray(nums []int) int {

	mx := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {

		sum = max(sum+nums[i], nums[i])
		mx = max(sum, mx)

	}

	return mx
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
