package main

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(robberOne(0, n-1, nums), robberOne(1, n, nums))
}

func robberOne(i int, n int, nums []int) int {
	rob1, rob2, temp := 0, 0, 0
	for ; i < n; i++ {
		temp = max(rob2, rob1+nums[i])
		rob1 = rob2
		rob2 = temp
	}

	return rob2
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
