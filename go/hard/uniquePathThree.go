package main

import "fmt"

var pathNumber int
var max int

func uniquePathsIII(grid [][]int) int {

	pathNumber = 0
	max = 0

	m := len(grid)
	n := len(grid[0])
	max = m * n
	start := []int{0, 0}
	// the square count which in the traveling path
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == -1 {
				max -= 1
			} else if grid[i][j] == 1 {
				start = []int{i, j}
			}
		}
	}

	travel(&grid, start[0], start[1], [][]int{})

	return pathNumber
}

func travel(grid *[][]int, x int, y int, path [][]int) bool {

	if x < 0 || y < 0 || x >= len(*grid) || y >= len((*grid)[0]) || (*grid)[x][y] == -1 {
		return true
	}

	for _, v := range path {
		pX := v[0]
		pY := v[1]
		if x == pX && y == pY {
			return true
		}

	}
	path = append(path, []int{x, y})

	if (*grid)[x][y] == 2 && len(path) == max {
		pathNumber++
		return true
	}

	// left
	travel(grid, x, y-1, path)

	// top
	travel(grid, x-1, y, path)

	// right
	travel(grid, x, y+1, path)

	// bottom
	travel(grid, x+1, y, path)

	return true
}

func main() {
	// fmt.Println(uniquePathsIII(
	// 	[][]int{
	// 		{1, 0, 0, 0},
	// 		{0, 0, 0, 0},
	// 		{0, 0, 2, -1},
	// 	},
	// ))
	// fmt.Println(uniquePathsIII(
	// 	[][]int{
	// 		{0, 1},
	// 		{2, 0},
	// 	},
	// ))
	// fmt.Println(uniquePathsIII(
	// 	[][]int{
	// 		{1, 0, 0, 0},
	// 		{0, 0, 0, 0},
	// 		{0, 0, 0, 2},
	// 	},
	// ))
	fmt.Println(uniquePathsIII(
		[][]int{
			{1, 0, 2},
		},
	))
}
