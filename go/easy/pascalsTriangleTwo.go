package main

import "fmt"

func getRow(rowIndex int) []int {

	arr := [][]int{{1}}

	if rowIndex >= 1 {
		arr = append(arr, []int{1, 1})
	}

	for i := 2; i <= rowIndex; i++ {
		nextRow := []int{1}

		preRow := arr[len(arr)-1]

		for i := 0; i < len(preRow)-1; i++ {
			nextRow = append(nextRow, preRow[i]+preRow[i+1])
		}
		nextRow = append(nextRow, 1)
		arr = append(arr[1:], nextRow)
	}
	return arr[len(arr)-1]
}

func main() {

	fmt.Println(getRow(5))

}
