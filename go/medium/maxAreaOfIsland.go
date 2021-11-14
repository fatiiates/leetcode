package main

import "fmt"

var maxArea int
var localArea int
var visitedCells map[int]map[int]bool

func maxAreaOfIsland(grid [][]int) int {

	maxArea = 0
	localArea = 0
	visitedCells = make(map[int]map[int]bool)

	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			travel(&grid, i, j)
		}
	}

	return maxArea
}

func travel(i *[][]int, x int, y int) bool {

	if x < 0 || y < 0 || x >= len(*i) || y >= len((*i)[0]) {
		return true
	}

	if _, ok := visitedCells[x]; ok {
		if _, ok := visitedCells[x][y]; ok {
			return true
		} else {
			visitedCells[x][y] = true
		}
	} else {
		visitedCells[x] = make(map[int]bool)
		visitedCells[x][y] = true
	}

	isStartOfIsland := false
	if (*i)[x][y] == 1 {
		if localArea == 0 {
			isStartOfIsland = true
		}
		localArea += 1
	} else if localArea > 0 {
		return false
	}

	// left
	travel(i, x, y-1)

	// top
	travel(i, x-1, y)

	// right
	travel(i, x, y+1)

	// bottom
	travel(i, x+1, y)

	if isStartOfIsland {
		maxArea = max(maxArea, localArea)
		localArea = 0
	}

	return true
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {

	fmt.Println(maxAreaOfIsland([][]int{
		{0, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}))

}
