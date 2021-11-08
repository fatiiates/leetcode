package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {

	globalStr := s
	for {
		localStr := globalStr
		localStr = strings.Replace(localStr, "()", "", 1)
		localStr = strings.Replace(localStr, "[]", "", 1)
		localStr = strings.Replace(localStr, "{}", "", 1)

		if len(localStr) == len(globalStr) {
			break
		}
		globalStr = localStr
		if len(globalStr) == 0 {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(isValid("([])"))
	fmt.Println(isValid("{([])}"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))

}
