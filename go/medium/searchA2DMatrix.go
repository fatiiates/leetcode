package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])
	row := max(0, findRow(&matrix, 0, m-1, &target))
	return binarySearch(&matrix[row], 0, n-1, &target)
}

func findRow(nums *[][]int, s int, e int, t *int) int {
	p := (e - s) / 2
	// fmt.Println(s,e)
	if e < s {
		return e
	}

	if *t == (*nums)[s+p][0] {
		return s + p
	}

	if *t > (*nums)[s+p][0] {
		return findRow(nums, s+p+1, e, t)
	}

	return findRow(nums, s, s+p-1, t)
}

func binarySearch(nums *[]int, s int, e int, t *int) bool {
	p := (e - s) / 2
	if e < s {
		return false
	}

	if (*nums)[s+p] == *t {
		return true
	}

	if *t > (*nums)[s+p] {
		return binarySearch(nums, s+p+1, e, t)
	}

	return binarySearch(nums, s, s+p-1, t)
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {

	fmt.Println(searchMatrix([][]int{
		{-1},
		// {9, 11, 16, 20},
		// {23, 30, 34, 35},
	}, -1))
}
