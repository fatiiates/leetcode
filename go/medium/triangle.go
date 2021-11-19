package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 1 {
		return triangle[0][0]
	}
	mn := math.MaxInt64
	for i := 1; i < m; i++ {

		n := len(triangle[i])
		for j := 0; j < n; j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == n-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
			}

			if i == m-1 {
				mn = min(mn, triangle[i][j])
			}
		}

	}

	return mn
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {

	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{5, 6, 7},
		{4, 1, 8, 3},
	}))
}
