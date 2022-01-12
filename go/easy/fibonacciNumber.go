package main

import "fmt"

func fib(n int) int {

	if n == 0 {
		return 0
	}

	first := 0
	second := 1
	for i := n-1; i > 0; i-- {
		first, second = second, first+second
	}
	return second

}

func main() {
	fmt.Println(fib(0)) // 0
	fmt.Println(fib(1)) // 1
	fmt.Println(fib(2)) // 1
	fmt.Println(fib(3)) // 2
	fmt.Println(fib(4)) // 3
	fmt.Println(fib(5)) // 5
	fmt.Println(fib(6)) // 8
	fmt.Println(fib(7)) // 13
}
