package main

import "fmt"

func maxArea(height []int) int {
	amount := 0
	l := len(height) - 1
	for i, j := 0, l; i < j; {
		fmt.Println(i, j)
		x := j - i
		y := min(height[i], height[j])
		amount = max(amount, x*y)
		if height[i] < height[j] {
			tmp := i
			for k := i + 1; k < j; k++ {
				if height[k] > height[tmp] {
					i = k
					break
				}
			}
			if tmp == i {
				break
			}
		} else {
			tmp := j
			for k := j; k > i; k-- {
				if height[k] > height[tmp] {
					j = k
					break
				}
			}
			if tmp == j {
				break
			}
		}
	}

	return amount
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

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1, 3, 2, 5, 25, 24, 5}))
}

/*

0,6 - 6
1,6 - 15
3,6 - 15
3,5

*/