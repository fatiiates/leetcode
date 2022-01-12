package main

import "fmt"

func tribonacci(n int) int {
    m := make(map[int]int)
	m[0] = 0
	m[1] = 1
	m[2] = 1

	for i := 3; i <= n; i++ {
		m[i] = m[i-3] + m[i-2] + m[i-1]
	}
	return m[n]
}

func main() {
	fmt.Println(tribonacci(0)) // 0
	fmt.Println(tribonacci(1)) // 1
	fmt.Println(tribonacci(2)) // 1
	fmt.Println(tribonacci(3)) // 2
	fmt.Println(tribonacci(25)) // 1389537
}
