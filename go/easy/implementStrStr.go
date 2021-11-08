package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	}

	ln := len(needle)
	lh := len(haystack)
	l := lh - ln
	for i := 0; i <= l; i++ {
		if haystack[i] == needle[0] {
			ctrl := true
			for j := 1; j < ln; j++ {
				if haystack[i+j] != needle[j] {
					ctrl = false
					break
				}
			}
			if ctrl {
				return i
			}
		}
	}

	return -1
}

func main() {

	fmt.Println(strStr("hello", "ll"))

}
