package main

import "fmt"

var paths [][]int

func allPathsSourceTarget(graph [][]int) [][]int {

	l := len(graph)
	last := l - 1
	paths = [][]int{}
	travel(&graph, 0, last, []int{})

	return paths
}

func travel(graph *[][]int, n int, last int, path []int) {
	path = append(path, n)
	if n == last {
		arr := make([]int, len(path))
		copy(arr, path)
		paths = append(paths, arr)
	} else {
		l := len((*graph)[n])

		for i := 0; i < l; i++ {

			travel(graph, (*graph)[n][i], last, path)
		}
	}

}

func main() {

	fmt.Println(allPathsSourceTarget([][]int{
		{3, 1},
		{4, 6, 7, 2, 5},
		{4, 6, 3},
		{6, 4},
		{7, 6, 5},
		{6},
		{7},
		{},
	}))
}

// [[0,3,6,7],[0,3,4,5],
// [[0,3,6,7],[0,3,4,7],
