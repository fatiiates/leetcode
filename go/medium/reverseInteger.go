package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {

	s := 1
	if x < 0 {
		s = -1
		x *= s
	}
	nums := []int{}
	for ; x != 0; x /= 10 {
		nums = append(nums, x%10)
	}

	l := len(nums)
	maxInt := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}

	if s == 1 {

		if l == 10 {
			for i := 0; i < len(maxInt); i++ {
				if nums[i] > maxInt[i] {
					return 0
				} else if nums[i] < maxInt[i] {
					break
				}
			}
		}
	} else {
		if l == 10 {

			for i := 0; i < len(maxInt)-1; i++ {
				if nums[i] > maxInt[i] {
					return 0
				} else if nums[i] < maxInt[i] {
					break
				}
			}
			if nums[0] > maxInt[9]-1 {
				return 0
			}
		}
	}

	rx := 0

	for i, v := range nums {
		rx += int(math.Pow10(l-1-i)) * v
	}

	return s * rx
}

func main() {

	fmt.Println(reverse(1147483641))

}
