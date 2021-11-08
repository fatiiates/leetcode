package main

import "fmt"

func generate(numRows int) [][]int {

	arr := [][]int{}

	if numRows == 1 {
		arr = append(arr, []int{1})
	} else if numRows >= 2 {
		arr = append(arr, []int{1})
		arr = append(arr, []int{1, 1})
	}

	for i := 2; i < numRows; i++ {
		nextRow := []int{1}

		preRow := arr[len(arr)-1]

		for i := 0; i < len(preRow)-1; i++ {
			nextRow = append(nextRow, preRow[i]+preRow[i+1])
		}
		nextRow = append(nextRow, 1)
		arr = append(arr, nextRow)
	}
	return arr
}

func main() {

	fmt.Println(generate(5))

}
