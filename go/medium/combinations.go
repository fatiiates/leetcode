package main

import "fmt"

var global [][]int

func combine(n int, k int) [][]int {
	global = [][]int{}
	combination(1, n, k, []int{})
	return global
}

func combination(start int, n int, k int, local []int) {

	if len(local) == k {
		arr := make([]int, len(local))
		copy(arr, local)
		global = append(global, arr)
		return
	}

	for i := start; i <= n; i++ {
		local = append(local, i)
		combination(i+1, n, k, local)
		local = local[0 : len(local)-1]
	}

}

func main() {

	fmt.Println(combine(20, 10))
}
