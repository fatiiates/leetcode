package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	travel(&image, sr, sc, [][]int{}, image[sr][sc], newColor)

	return image
}

func travel(i *[][]int, x int, y int, path [][]int, oldColor int, newColor int) bool {

	if x < 0 || y < 0 || x >= len(*i) || y >= len((*i)[0]) {
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

	if (*i)[x][y] == oldColor {
		(*i)[x][y] = newColor
	} else {
		return false
	}

	// left
	travel(i, x, y-1, path, oldColor, newColor)

	// top
	travel(i, x-1, y, path, oldColor, newColor)

	// right
	travel(i, x, y+1, path, oldColor, newColor)

	// bottom
	travel(i, x+1, y, path, oldColor, newColor)

	return true
}

func main() {

	fmt.Println(floodFill([][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}, 1, 1, 2))

}
