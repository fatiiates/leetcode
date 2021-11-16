package main

import (
	"fmt"
)

func orangesRotting(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	freshs := 0

	var q [][2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			} else if grid[i][j] == 1 {
				freshs += 1
			}
		}
	}

	mn := 0

	for len(q) > 0 {
		l := len(q)

		for i := 0; i < l; i++ {
			c := q[0]
			q = q[1:]
			n := neighbors(c[0], c[1], n, m)

			for _, p := range n {
				if grid[p[0]][p[1]] == 1 {
					grid[p[0]][p[1]] = 2
					freshs -= 1
					q = append(q, [2]int{p[0], p[1]})
					q = append(q, [2]int{p[0], p[1]})
				}
			}
		}
		if len(q) > 0 {
			mn += 1
		}
	}

	if freshs > 0 {
		return -1
	}

	return mn
}

func neighbors(row int, col int, m int, n int) [][]int {
	arr := [][]int{}
	if row+1 < n {
		arr = append(arr, []int{row + 1, col})
	}
	if row-1 >= 0 {
		arr = append(arr, []int{row - 1, col})
	}
	if col+1 < m {
		arr = append(arr, []int{row, col + 1})
	}
	if col-1 >= 0 {
		arr = append(arr, []int{row, col - 1})
	}
	return arr
}

func main() {

	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}))
}
