package main

import "fmt"

func findKthNumber(m int, n int, k int) int {
	s := 1
	e := m * n
	for s < e {
		p := (s + e) / 2
		c := 0
		for i, j := 1, n; i <= m; i++ {
			for j >= 1 && i*j > p {
				j--
			}
			c += j
		}
		if c >= k {
			e = p
		} else {
			s = p + 1
		}
	}
	return s
}

func main() {
	fmt.Println(findKthNumber(12, 68, 5))
}
