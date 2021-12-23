package main

import "fmt"

var mod int = 1e9 + 7

func nthMagicalNumber(n int, a int, b int) int {

	l := a / gcd(a, b) * b

	var lo int64 = 0
	var hi int64 = int64(n) * int64(min(a, b))
	for lo < hi {
		var mi int64 = lo + (hi-lo)/2

		if (mi/int64(a) + mi/int64(b) - mi/int64(l)) < int64(n) {
			lo = mi + 1
		} else {
			hi = mi
		}
	}

	return int(lo % int64(mod))
}

func gcd(x int, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {

	fmt.Println(nthMagicalNumber(1000000000, 40000, 40000))

}
