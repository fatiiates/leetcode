package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {

	arr := make([]string, numRows)
	mod := numRows*2 - 2
	if mod <= 0 {
		mod = 1
	}
	diff := mod - numRows
	for i := 0; i < len(s); i++ {
		row := i % mod
		if row >= numRows {
			row = numRows - row + diff
		}
		arr[row] += string(s[i])
	}
	return strings.Join(arr, "")
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}
