package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {

	x := strings.Fields(s)

	return len(x[len(x)-1])
}
