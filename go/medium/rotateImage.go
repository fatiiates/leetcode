package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix)
	// tranpose matrix
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	fmt.Println(matrix)
	// reverse all matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

func main() {
	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// rotate(matrix)
	// fmt.Println(matrix)

	// matrix := [][]int{
	// 	{5, 1, 9, 11},
	// 	{2, 4, 8, 10},
	// 	{13, 3, 6, 7},
	// 	{15, 14, 12, 16},
	// }
	// rotate(matrix)
	// fmt.Println(matrix)
	
	matrix := [][]int{
		{1,2,3,4,5},
		{6,7,8,9,10},
		{11,12,13,14,15},
		{16,17,18,19,20},
		{21,22,23,24,25},
	}
	rotate(matrix)
	fmt.Println(matrix)

}
