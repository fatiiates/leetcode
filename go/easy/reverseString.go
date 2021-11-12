package main

import "fmt"

func reverseString(s []byte) {
	fmt.Println(s)

	l := len(s)
	if l > 1 {

		for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	fmt.Println(s)
}

func main() {

	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})

}
