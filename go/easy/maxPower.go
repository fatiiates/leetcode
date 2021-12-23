package main

import "fmt"

func maxPower(s string) int {
	c := rune(s[0])
	mx := 0
	localMx := 0
	for _, v := range s {
		if v == c {
			localMx++
		} else {
			mx = max(localMx, mx)
			localMx = 1
		}
		c = v
	}
	return max(localMx, mx)
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	fmt.Println(maxPower("leetcode"))
}
