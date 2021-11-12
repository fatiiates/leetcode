package main

import "fmt"

func reverseWords(s string) string {

	pre := 0
	i := 0
	l := len(s)
	newS := ""
	for ; i < l; i++ {

		if s[i] == ' ' {

			newS += reverseString([]byte(s[pre:i])) + " "
			pre = i + 1
		}
	}

	newS += reverseString([]byte(s[pre:]))

	return newS

}

func reverseString(s []byte) string {

	l := len(s)
	if l > 1 {

		for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	return string(s)
}

func main() {

	fmt.Println(reverseWords("Let's take LeetCode contest "))

}
