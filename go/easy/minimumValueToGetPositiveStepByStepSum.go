package main

import "fmt"

func minStartValue(nums []int) int {

	l := len(nums)
	m := 0
	t := 0
	for i := 0; i < l; i++ {
		t += nums[i]
		if t < 0 {
			m = min(m, t)
		}
	}
	return abs(m) + 1

}

func min(n1 int, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func main() {

	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))

}
