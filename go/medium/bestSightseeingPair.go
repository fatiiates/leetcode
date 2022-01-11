package main

import "fmt"

// A[i] + i (A[j] - j)
func maxScoreSightseeingPair(values []int) int {
	mx := values[0]
	ans := 0
	l := len(values)
	for i := 1; i < l; i++ {
		if ans < mx+values[i]-i {
			ans = mx+values[i]-i
		}
		if mx < values[i]+i {
			mx = values[i]+i
		}
	}
	return ans
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
}
