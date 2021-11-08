package main

func maxSubArray(nums []int) int {

	max := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {

		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if max < sum {
			max = sum
		}

	}

	return max
}
