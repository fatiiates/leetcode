package main

import (
	"fmt"
	"math"
)

var bad int = 4

func firstBadVersion(n int) int {
	return search(0, n)
}

func isBadVersion(version int) bool {
	if version > bad {
		return true
	}
	return false
}

func search(start int, last int) int {

	med := int(math.Ceil(float64((last - start)) / 2.0))

	res := isBadVersion(start + med)

	if !res && last == start {
		return start + 1
	}
	if res {
		return search(start, last-med)
	} else {
		return search(start+med, last)
	}

}

func main() {

	fmt.Println(firstBadVersion(5))
}
