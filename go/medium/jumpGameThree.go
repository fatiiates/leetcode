package main

import "fmt"

var visited map[int]bool
var l int

func canReach(arr []int, start int) bool {
	l = len(arr)
	visited = make(map[int]bool)
	return travel(&arr, start)

}

func travel(arr *[]int, i int) bool {
	if i >= 0 && i < l {
		if (*arr)[i] == 0 {
			return true
		}
		visited[i] = true

		mx := i + (*arr)[i]
		mn := i - (*arr)[i]

		ctrl := false
		if _, ok := visited[mn]; !ok && mn >= 0 {
			ctrl = travel(arr, mn)
		}
		if _, ok := visited[mx]; !ok && !ctrl && mx < l {
			ctrl = travel(arr, mx)
		}
		return ctrl
	} else {
		return false
	}
}

func main() {

	fmt.Println(canReach([]int{2, 1, 2, 9, 3, 4}, 5))
}
