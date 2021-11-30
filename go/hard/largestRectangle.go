package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {

	m := len(matrix)
	mx := 0

	if m <= 0 {
		return 0
	}

	n := len(matrix[0])

	row := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[j] = 0
			} else {
				row[j] += 1
			}
		}

		mx = max(mx, hist(row))
	}

	return mx
}

func hist(heights []int) int {
	st := []int{}
	mx := 0
	heights = append(heights, 0)
	for i, v := range heights {
		l := len(st) - 1
		for l >= 0 && heights[st[l]] >= v {
			height := heights[pop(&st)]
			width := 0
			l = len(st) - 1
			if l < 0 {
				width = i
			} else {
				width = i - st[l] - 1
			}
			mx = max(mx, height*width)
		}
		st = append(st, i)
	}

	return mx
}

func pop(arr *[]int) int {
	l := len(*arr)
	last := (*arr)[l-1]
	*arr = (*arr)[:l-1]
	return last
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {

	fmt.Println(maximalRectangle([][]byte{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}))

}
