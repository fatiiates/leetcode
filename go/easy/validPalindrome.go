package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {

	re := regexp.MustCompile(`[a-zA-Z0-9]+`)

	if s == "" {
		return true
	}
	submatchall := re.FindAllString(s, -1)
	s = strings.ToLower(strings.Join(submatchall, ""))

	l := len(s)
	for i, v := range s {
		if byte(v) != s[l-i-1] {
			return false
		}
	}

	return true

}

func main() {

	fmt.Println(isPalindrome("0P"))

}
