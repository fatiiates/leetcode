package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {

	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if (n | (n - 1)) == 2*n-1 {
		return true
	} else {
		return false
	}
}

func main() {

	fmt.Println(isPowerOfTwo(5))
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(0))
	fmt.Println(isPowerOfTwo(16))
}
