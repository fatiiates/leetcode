package main

import (
	"fmt"
)

var p []int

func largestComponentSize(nums []int) int {

	p = []int{}

	l := len(nums)

	mx := nums[0]
	for i := 1; i < l; i++ {
		mx = max(mx, nums[i])
	}

	for i := 0; i <= mx; i++ {
		p = append(p, i)
	}

	for _, v := range nums {
		for i := 2; i*i <= v; i++ {
			if v%i == 0 {
				union(v, i)
				union(v, v/i)
			}
		}
	}

	var m map[int]int = make(map[int]int)
	res := 1
	for i := 0; i < l; i++ {
		f := find(nums[i])
		m[f]++
		res = max(res, m[f])
	}

	return res
}

func max(n1 int, n2 int) int {
	if n2 > n1 {
		return n2
	}
	return n1
}

func union(x int, y int) {
	x = find(x)
	y = find(y)
	if x == y {
		return
	}
	p[x] = p[y]
}

func find(x int) int {
	if p[x] == x {
		return x
	}
	p[x] = find(p[x])
	return p[x]
}

func main() {
	fmt.Println(largestComponentSize([]int{20, 50, 9, 63}))
}
