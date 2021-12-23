package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	r := make(map[int]map[byte]bool)
	c := make(map[int]map[byte]bool)
	b := make(map[int]map[byte]bool)
	for i := 0; i < 9; i++ {
		r[i] = make(map[byte]bool)
		c[i] = make(map[byte]bool)
		b[i] = make(map[byte]bool)
	}

	for i, row := range board {
		for j, v := range row {
			if v != '.' {
				if _, ok := r[i][v]; ok {
					return false
				}

				if _, ok := c[j][v]; ok {
					return false
				}
				bi := i/3*3 + j/3
				if _, ok := b[bi][v]; ok {
					return false
				}

				r[i][v] = true
				c[j][v] = true
				b[bi][v] = true
			}
		}
	}

	return true
}

func main() {

	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))

}
