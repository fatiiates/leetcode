package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {

	m := len(mat)
	n := len(mat[0])
	if m*n != r*c || (r == m && c == n) {
		return mat
	}

	arr := [][]int{}

	for i := 0; i < r; i++ {
		arr = append(arr, make([]int, c))
	}
	k := 0

	for _, row := range mat {
		for _, col := range row {
			arr[k/c][k%c] = col
			k++
		}
	}

	return arr
}

func main() {

	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))

}
